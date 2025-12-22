package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type solution struct {
	input      []string
	ans1, ans2 int
}

func buildSolution(r io.Reader) *solution {
	lines, _ := utils.LinesFromReader(r)
	return &solution{
		input: lines,
		ans1:  0,
		ans2:  0,
	}
}

func score(r rune) int {
	if r >= 'a' {
		return int(r - 'a' + 1)
	}
	return int(r - 'A' + 27)
}

func (s *solution) run1() {
	m := [128]int{}
	for _, ruck := range s.input {
		half := len(ruck) / 2
		for i, r := range ruck {
			if i < half { // first harf
				m[r]++ // track freq
			} else {
				if m[r] > 0 {
					s.ans1 += score(r)
					m = [128]int{}
					break
				}
			}
		}
	}
}

func (s *solution) findCommon(i int) rune {
	var m [128]int
	for j := range 3 {
		var local [128]int
		for _, r := range s.input[i+j] {
			if local[r] > 0 {
				continue
			}
			local[r]++
			m[r]++
		}
		// for i, v := range local {
		// 	if v > 0 {
		// 		m[i]++
		// 	}
		// }
	}
	for i, v := range m {
		if v == 3 {
			return rune(i)
		}
	}
	return 0
}

func (s *solution) run2() {
	for i := 0; i < len(s.input); i += 3 {
		s.ans2 += score(s.findCommon(i))
	}
}

func (s *solution) res1() int {
	return s.ans1
}

func (s *solution) res2() int {
	return s.ans2
}

func part1(r io.Reader) int {
	s := buildSolution(r)
	s.run1()
	return s.res1()
}

func part2(r io.Reader) int {
	s := buildSolution(r)
	s.run2()
	return s.res2()
}

func main() {
	Input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("fail open input.txt %v", err)
	}
	start := time.Now()
	result := part1(Input)
	elapsed := time.Since(start)
	fmt.Printf("p1 res 🙆-> %d (Time taken: %s)\n", result, elapsed)
	start = time.Now()
	result = part2(Input)
	elapsed = time.Since(start)
	fmt.Printf("p2 res 🙆-> %d (Time taken: %s)\n", result, elapsed)
}
