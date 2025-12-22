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

type pair struct {
	l, r int
}

func intersect(p1, p2 pair) pair {
	l, r := max(p1.l, p2.l), min(p1.r, p2.r)
	return pair{l, r}
}

// func contain(p1, p2 pair) bool {
// 	// if l <= r { // overlap
// 	inters :=
// 	return inters == p1 || inters == p2
// 	// if inters == p1 || inters == p2 {
// 	// 		return true
// 	// 	}
// 	// }
// 	// return false
// }

func (s *solution) run1() {
	var a, b, c, d int
	for _, str := range s.input {
		_, err := fmt.Sscanf(str, "%d-%d,%d-%d", &a, &b, &c, &d)
		if err != nil {
			panic(err)
		}
		p1, p2 := pair{a, b}, pair{c, d}
		inter := intersect(p1, p2)
		if inter.l <= inter.r { // if intersect exist
			s.ans2++
		}
		if inter == p1 || inter == p2 {
			s.ans1++
		}
	}
}

func (s *solution) run2() {
	s.run1()
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
