package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func (s *solution) isValid(report []int) bool {
	initDiff := report[1] - report[0]
	for i := 1; i < len(report); i++ {
		curdiff := report[i] - report[i-1]
		if curdiff*initDiff <= 0 || utils.Abs(curdiff) > 3 {
			return false
		}
	}
	return true
}

func (s *solution) isValid2(report []int) bool {
	// generate a new slice with one less item
	newReport := make([]int, 0, len(report)-1)
	for i := range report {
		for j, v := range report {
			if i == j {
				continue
			}
			newReport = append(newReport, v)
		}
		if s.isValid(newReport) {
			return true
		}
		newReport = newReport[:0]
	}

	return false
}

func (s *solution) run1() {
	for _, report := range s.reports {
		if s.isValid(report) {
			s.ans++
		}
	}
}

func (s *solution) run2() {
	for _, report := range s.reports {
		if s.isValid(report) || s.isValid2(report) {
			s.ans++
		}
	}
}

func (s *solution) res() int {
	return s.ans
}

type solution struct {
	ans     int
	reports [][]int
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}
	var reports [][]int
	for _, line := range lines {
		ints := utils.IntsFromString(line)
		reports = append(reports, ints)
	}

	return &solution{
		ans:     0,
		reports: reports,
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
	start := time.Now()
	result := part1(Input)
	elapsed := time.Since(start)
	fmt.Printf("p1 res 🙆-> %d (Time taken: %s)\n", result, elapsed)
	start = time.Now()
	result = part2(Input)
	elapsed = time.Since(start)
	fmt.Printf("p2 res 🙆-> %d (Time taken: %s)\n", result, elapsed)
}
