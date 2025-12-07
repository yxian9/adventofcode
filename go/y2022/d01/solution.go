package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
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

func (s *solution) run1() {
	var (
		cals []int
		sum  int
	)

	for _, cal := range s.input {
		if cal == "" {
			cals = append(cals, sum)
			sum = 0
		}
		i, _ := strconv.Atoi(cal)
		sum += i
	}

	sort.Ints(cals)

	s.ans1 = cals[len(cals)-1]

	for i := range 3 {
		s.ans2 += cals[len(cals)-i-1]
	}

	// fmt.Printf("%#v\n", s.input)
	// var (
	// 	sum int
	// )
	// for _, item := range s.input {
	// 	i, _ := strconv.Atoi(item)
	// 	sum += i
	// 	if i == 0 {
	// 		s.ans1 = max(sum, s.ans1)
	// 		sum = 0
	// 	}
	// }
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
