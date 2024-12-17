package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
)

type solution struct {
	l1  []int
	l2  []int
	ans int
}

func (s *solution) run1() {
	slices.Sort(s.l1)
	slices.Sort(s.l2)
	for i := range s.l1 {
		s.ans += utils.Abs(s.l1[i] - s.l2[i])
	}
}

func (s *solution) run2() {
	counts := make(map[int]int)
	for _, num := range s.l1 {
		counts[num]++
	}

	for _, num := range s.l2 {
		s.ans += num * counts[num]
	}
}

func (s *solution) res() int {
	return s.ans
}

func buildSolution(r io.Reader) solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}
	var (
		l1 = make([]int, len(lines))
		l2 = make([]int, len(lines))
	)
	for i, line := range lines {
		ints := utils.IntsFromString(line)
		l1[i], l2[i] = ints[0], ints[1]
	}

	return solution{l1, l2, 0}
}

func Part1(r io.Reader) int {
	solution := buildSolution(r)
	solution.run1()
	return solution.res()
}

func Part2(r io.Reader) int {
	solution := buildSolution(r)
	solution.run2()
	return solution.res()
}

func main() {
	arg := os.Args[1]
	fmt.Println("Running part", arg)
	switch arg {
	case "1":
		fmt.Println("p1 res 🙆-> ", Part1(os.Stdin))
	case "2":
		fmt.Println("p2 res 🙆-> ", Part2(os.Stdin))
	}
}
