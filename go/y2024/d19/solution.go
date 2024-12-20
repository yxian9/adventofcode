package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type solution struct {
	pattens        map[string]bool
	designs        []string
	memo, itemLens map[int]bool
	memo2          map[int]int
	ans            int
}

func (s *solution) dfs(start int, design string) bool {
	if start == len(design) {
		return true
	}
	if res, ok := s.memo[start]; ok {
		return res
	}
	res := false
	for pattern := range s.pattens {
		end := start + len(pattern)
		if end > len(design) {
			continue
		}
		cur_part := design[start:end]
		if s.pattens[cur_part] {
			res = s.dfs(end, design)
			if res {
				break
			}
		}
	}
	s.memo[start] = res

	return res
}

func (s *solution) dfs2(start int, design string) int {
	if start == len(design) {
		return 1
	}
	if res, ok := s.memo2[start]; ok {
		return res
	}
	res := 0
	for patternLen := range s.itemLens {
		end := start + patternLen
		if end > len(design) {
			continue
		}
		cur_part := design[start:end]
		if s.pattens[cur_part] {
			res += s.dfs2(end, design)
		}
	}
	s.memo2[start] = res

	return res
}

func (s *solution) run1() {
	for _, design := range s.designs {
		s.memo = make(map[int]bool)
		if s.dfs(0, design) {
			s.ans++
		}
	}
}

func (s *solution) run2() {
	for _, design := range s.designs {
		s.memo2 = make(map[int]int)
		s.ans += s.dfs2(0, design)
	}
}

func (s *solution) res() int {
	return s.ans
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}
	designs := []string{}
	itemLens := map[int]bool{}
	patterns := map[string]bool{}

	for i, line := range lines {
		if i == 0 {
			items := strings.Split(line, ", ")
			for _, v := range items {
				itemLens[len(v)] = true
				patterns[v] = true
			}
		}
		if i == 1 {
			continue
		}
		designs = append(designs, line)
	}

	return &solution{
		designs:  designs,
		pattens:  patterns,
		itemLens: itemLens,
		ans:      0,
	}
}

func part1(r io.Reader) int {
	s := buildSolution(r)
	s.run1()
	return s.res()
}

func part2(r io.Reader) int {
	s := buildSolution(r)
	s.run2()
	return s.res()
}

func main() {
	Input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("fail open input.txt %v", err)
	}
	defer Input.Close()
	fmt.Println("p1 res 🙆-> ", part1(Input))
	fmt.Println("p2 res 🙆-> ", part2(Input))
}
