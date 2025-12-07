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

func findJolt(bank string, length int) (output int) {
	idx := -1
	for i := length; i > 0; i-- {
		var jolts int
		jolts, idx = findMax(bank, idx+1, len(bank)-i+1)
		output = output*10 + jolts
	}
	return output
}

func findMax(bank string, start, end int) (res, idx int) {
	for i := start; i < end; i++ {
		v := int(bank[i] - '0')
		if v > res { // can not be >=, need keep 2nd largest as 2nd
			res = v
			idx = i
		}
	}
	return res, idx
}

func (s *solution) run1() {
	for _, bank := range s.input {
		s.ans1 += findJolt(bank, 2)
	}
}

func (s *solution) run2() {
	for _, bank := range s.input {
		s.ans2 += findJolt(bank, 12)
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
