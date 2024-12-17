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
	//go:embed cookie.txt
	cookie string
)

func main() {
	curPath, _ := os.Getwd()
	day, year, overwrite := ParseFlags()
	gen, err := NewGenerator(day, year, curPath, overwrite)
	// fmt.Printf("%#v \n", gen)
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
	flag.IntVar(&year, "y", today.Year(), "AOC year")
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

// A Generator creates a directory with all contents required to kickstart
// a solution to a puzzle in the Advent of Code calendar.
type Generator struct {
	// The day and year to generate code for.
	day, year int

	// The directory where all Advent of Code solutions are stored.
	workdir string

	// Whether to overwrite existing files.
	overwrite bool

	// Path to scaffolded directory.
	packageDir string
}

// NewGenerator builds a generator for the given date and author. If overwrite
// is true, the generator will overwrite existing files.
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

// Initialize validates gen's parameters and pre-computes useful values.
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

	gen.setPackageDir()

	return nil
}

func (gen *Generator) setPackageDir() {
	gen.packageDir = filepath.Join(
		gen.workdir,
		fmt.Sprintf("y%04d", gen.year),
		fmt.Sprintf("d%02d", gen.day),
	)
}

// Run uses gen to build scaffolding for an Advent of Code solution.
func (gen *Generator) Run() error {
	if err := gen.CreatePackage(); err != nil {
		return fmt.Errorf("creating package: %w", err)
	}
	if err := gen.WriteCode(); err != nil {
		return fmt.Errorf("writing code: %w", err)
	}
	if err := gen.DownloadInput(); err != nil {
		return fmt.Errorf("downloading input: %w", err)
	}
	return nil
}

// CreatePackage creates a directory/package to put scaffolding into.
func (gen *Generator) CreatePackage() error {
	err := os.MkdirAll(gen.packageDir, 0755)
	if err != nil {
		return fmt.Errorf("creating directory %q: %w", gen.packageDir, err)
	}

	fmt.Printf("🏗️  Scaffolding package: %s\n", gen.packageDir)
	return nil
}

// WriteCode builds Go scaffolding for implementing, testing, and benchmarking
// solutions to Advent of Code problems.
func (gen *Generator) WriteCode() error {
	if err := gen.copyTemplatetoFile([]byte(solutionTemplate), "solution.go"); err != nil {
		return fmt.Errorf("creating %q: %w", "solution.go", err)
	}
	if err := gen.copyTemplatetoFile([]byte(solutionTestTemplate), "solution_test.go"); err != nil {
		return fmt.Errorf("creating %q: %w", "solution_test.go", err)
	}
	if err := gen.copyTemplatetoFile([]byte(testInput), "test1.txt"); err != nil {
		return fmt.Errorf("creating %q: %w", "test1.txt", err)
	}
	return nil
}

func (gen *Generator) copyTemplatetoFile(contents []byte, filename string) error {
	path := filepath.Join(gen.packageDir, filename)

	if fileExists(path) && !gen.overwrite {
		fmt.Printf("  👉 Skipping existing file %s.\n", path)
		return nil
	}

	if err := os.WriteFile(path, contents, os.FileMode(0644)); err != nil {
		return fmt.Errorf("writing %q: %w", path, err)
	}

	return nil
}

// fileExists checks whether filename exists.
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

// DownloadInput fetches the Advent of Code's daily input and writes it to a
// testdata directory.
func (gen *Generator) DownloadInput() error {
	path := filepath.Join(gen.packageDir, "input.txt")
	if fileExists(path) && !gen.overwrite {
		fmt.Println("  👉 Skipping input download; file already exists.")
		return nil
	}

	if cookie == "" {
		fmt.Println("  👉 Skipping input download; no session cookie provided.")
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

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("creating directory %q: %w", filepath.Dir(path), err)
	}

	err = os.WriteFile(path, input, 0644)
	if err != nil {
		return fmt.Errorf("writing input to file %q: %w", path, err)
	}

	fmt.Printf("  👉 Downloaded input.\n")

	return nil
}
