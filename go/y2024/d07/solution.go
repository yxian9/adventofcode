package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func (s *solution) run1() {
	for _, ints := range s.input {
		if s.valid(0, ints[1], ints[0], ints[2:], s.ops) {
			s.ans += ints[0]
		}
	}
}

func (s *solution) valid(idx, res, target int, ints []int, ops []op) bool {
	if idx == len(ints) {
		return res == target
	}
	if res > target {
		return false
	}

	for _, op := range ops {
		valid := s.valid(idx+1, op(res, ints[idx]), target, ints, ops)
		if valid {
			return true
		}
	}
	return false
}

func (s *solution) run2() {
	for _, ints := range s.input {
		if s.valid(0, ints[1], ints[0], ints[2:], s.ops2) {
			s.ans += ints[0]
		}
	}
}

func (s *solution) res() int {
	return s.ans
}

func (s *solution) res2() int {
	return s.ans
}

type op func(int, int) int

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	var input [][]int
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}
	for _, line := range lines {
		input = append(input, utils.IntsFromString(line))
	}
	ops := []op{
		func(a, b int) int { return a + b },
		func(a, b int) int { return a * b },
	}
	ops2 := []op{
		func(a, b int) int { return a + b },
		func(a, b int) int { return a * b },
		// func(a, b int) int {
		// 	factor := 1
		// 	if factor <= b {
		// 		factor *= 10
		// 	}
		// 	return a*factor + b
		// },
		func(a, b int) int {
			// if b == 0 {
			// 	return a * 10
			// }
			fac := 10
			for fac <= b {
				fac *= 10
			}
			return a*fac + b
		},
	}

	return &solution{
		input: input,
		ans:   0,
		ops:   ops,
		ops2:  ops2,
	}
}

type solution struct {
	input [][]int
	ans   int
	ops   []op
	ops2  []op
}

func part1(r io.Reader) int {
	s := buildSolution(r)
	s.run1()
	return s.res()
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
