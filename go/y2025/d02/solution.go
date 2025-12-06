package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"maps"
	"os"
	"time"
)

type solution struct {
	input      string
	ans1, ans2 int
}

func buildSolution(r io.Reader) *solution {
	line, _ := utils.ByteSFromReader(r)
	return &solution{
		input: string(line),
		ans1:  0,
		ans2:  0,
	}
}

func (s *solution) run1() {
	ranges := parseInput(s.input)
	for _, r := range ranges {
		if r.len%2 != 0 {
			continue
		}
		for n := range r.InvalidIDs(r.len / 2) {
			s.ans1 += n
		}
	}
}

func (s *solution) run2() {
	ranges := parseInput(s.input)
	invalids := map[int]struct{}{}

	for _, r := range ranges {
		for chunkLen := 1; chunkLen <= r.len/2; chunkLen++ {

			// only divisors
			if r.len%chunkLen != 0 {
				continue
			}

			maps.Copy(invalids, r.InvalidIDs(chunkLen))
		}
	}

	for n := range invalids {
		s.ans2 += n
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
