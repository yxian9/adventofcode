package main

import (
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var (
	//go:embed templ/test1.txt
	testInput string
	//go:embed templ/solution.go
	solutionTemplate string
	//go:embed templ/solution_test.go
	solutionTestTemplate string
	//go:embed templ/solution.py
	pySolutionTemplate string
	//go:embed templ/test_solution.py
	pyTestTemplate string
	//go:embed cookie.txt
	cookie string
)

type LangTarget struct {
	Name      string
	Dir       string
	Templates []FileTemplate
}

type FileTemplate struct {
	Filename string
	Content  []byte
}

func buildLangTargets(workdir string, year, day int) []LangTarget {
	dayDir := filepath.Join(fmt.Sprintf("y%04d", year), fmt.Sprintf("d%02d", day))
	return []LangTarget{
		{
			Name: "golang",
			Dir:  filepath.Join(workdir, "golang", dayDir),
			Templates: []FileTemplate{
				{"solution.go", []byte(solutionTemplate)},
				{"solution_test.go", []byte(solutionTestTemplate)},
				{"test1.txt", []byte(testInput)},
			},
		},
		{
			Name: "python",
			Dir:  filepath.Join(workdir, "python", dayDir),
			Templates: []FileTemplate{
				{"solution.py", []byte(pySolutionTemplate)},
				{"test_solution.py", []byte(pyTestTemplate)},
				{"test1.txt", []byte(testInput)},
			},
		},
	}
}

func main() {
	curPath, _ := os.Getwd()
	day, year, overwrite := ParseFlags()
	gen, err := NewGenerator(day, year, curPath, overwrite)
	if err != nil {
		log.Fatalf("making code generator: %v", err)
	}

	if err := gen.Run(); err != nil {
		log.Fatalf("building scaffolding: %v", err)
	}

	fmt.Println("🎅🏻 Merry coding!")
}

func ParseFlags() (day, year int, overwrite bool) {
	today := time.Now()
	flag.IntVar(&day, "d", today.Day(), "day number to fetch, 1-25")
	flag.IntVar(&year, "y", 2025, "AOC year")
	flag.BoolVar(&overwrite, "w", false, "overwrite?")
	flag.Parse()

	if day > 25 || day < 1 {
		log.Fatalf("day out of range: %d", day)
	}

	if year < 2015 {
		log.Fatalf("year is before 2015: %d", year)
	}

	if cookie == "" {
		log.Fatalf("no session cookie set on flag or env var (AOC_SESSION_COOKIE)")
	}

	return day, year, overwrite
}

type Generator struct {
	day, year int
	workdir   string
	overwrite bool
	langs     []LangTarget
}

func NewGenerator(day, year int, workdir string, overwrite bool) (*Generator, error) {
	gen := &Generator{
		day:       day,
		year:      year,
		workdir:   workdir,
		overwrite: overwrite,
	}

	if err := gen.Initialize(); err != nil {
		return nil, fmt.Errorf("failed initialization: %w", err)
	}

	return gen, nil
}

func (gen *Generator) Initialize() error {
	if gen.day <= 0 || gen.day > 25 {
		return fmt.Errorf("invalid day: %d", gen.day)
	}
	if gen.year <= 0 {
		return fmt.Errorf("invalid year: %d", gen.year)
	}
	if gen.workdir == "" {
		return errors.New("working directory unknown")
	}

	gen.langs = buildLangTargets(gen.workdir, gen.year, gen.day)

	return nil
}

func (gen *Generator) Run() error {
	for _, lang := range gen.langs {
		if err := gen.scaffold(lang); err != nil {
			return fmt.Errorf("scaffolding %s: %w", lang.Name, err)
		}
	}
	if err := gen.DownloadInput(); err != nil {
		return fmt.Errorf("downloading input: %w", err)
	}
	return nil
}

func (gen *Generator) scaffold(lang LangTarget) error {
	if err := os.MkdirAll(lang.Dir, 0755); err != nil {
		return fmt.Errorf("creating directory %q: %w", lang.Dir, err)
	}
	fmt.Printf("🏗️  Scaffolding %s: %s\n", lang.Name, lang.Dir)

	for _, tmpl := range lang.Templates {
		path := filepath.Join(lang.Dir, tmpl.Filename)
		if fileExists(path) && !gen.overwrite {
			fmt.Printf("  👉 Skipping existing file %s.\n", path)
			continue
		}
		if err := os.WriteFile(path, tmpl.Content, 0644); err != nil {
			return fmt.Errorf("writing %q: %w", path, err)
		}
	}
	return nil
}

func (gen *Generator) DownloadInput() error {
	if cookie == "" {
		fmt.Println("  👉 Skipping input download; no session cookie provided.")
		return nil
	}

	allExist := true
	for _, lang := range gen.langs {
		path := filepath.Join(lang.Dir, "input.txt")
		if !fileExists(path) || gen.overwrite {
			allExist = false
			break
		}
	}
	if allExist {
		fmt.Println("  👉 Skipping input download; files already exist.")
		return nil
	}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", gen.year, gen.day)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("preparing GET request to %q: %w", url, err)
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: cookie})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("sending GET request to %q: %w", url, err)
	}
	defer resp.Body.Close()

	input, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response from %q: %w", url, err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("adventofcode.com responded with %d: %s", resp.StatusCode, input)
	}

	for _, lang := range gen.langs {
		path := filepath.Join(lang.Dir, "input.txt")
		if fileExists(path) && !gen.overwrite {
			fmt.Printf("  👉 Skipping existing file %s.\n", path)
			continue
		}
		if err := os.WriteFile(path, input, 0644); err != nil {
			return fmt.Errorf("writing input to %q: %w", path, err)
		}
	}

	fmt.Printf("  👉 Downloaded input.\n")
	return nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		fmt.Println(filename)
		panic(err)
	}
	return !info.IsDir()
}
