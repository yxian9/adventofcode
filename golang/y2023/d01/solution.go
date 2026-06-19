package main

import (
	"adventofcode/h"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var digits = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

type solution struct {
	input      []string
	ans1, ans2 int
}

func buildSolution(r io.Reader) *solution {
	lines, _ := h.LinesFromReader(r)
	return &solution{
		input: lines,
		ans1:  0,
		ans2:  0,
	}
}

func is_dig(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}

func l_r(s string) (l, r int) {
	for _, r := range s {
		if is_dig(r) {
			l = int(r - '0')
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if is_dig(rune(s[i])) {
			r = int(s[i] - '0')
			break
		}
	}
	return
}

var step = [...]int{3, 4, 5}

func l2(s string) (l int) {

	for i := 0; i < len(s); i++ {
		if is_dig(rune(s[i])) {
			return int(s[i] - '0')
		}

		for _, add := range step {
			end := i + add
			if end >= len(s) {
				break
			}
			subs := s[i : i+add]
			if d, ok := digits[subs]; ok {
				return d
			}
		}
	}
	return 0
}

func r2(s string) (r int) {
	for i := len(s) - 1; i >= 0; i-- {
		if is_dig(rune(s[i])) {
			return int(s[i] - '0')
		}
		for _, add := range step {
			end := i + 1 - add
			if end < 0 {
				break
			}
			subs := s[end : i+1]
			if d, ok := digits[subs]; ok {
				return d
			}
		}
	}
	return
}

func (s *solution) run1() {
	for _, line := range s.input {
		l, r := l_r(line)
		s.ans1 += r
		s.ans1 += 10 * l
	}
}

func (s *solution) run2() {
	for _, line := range s.input {
		l := l2(line)
		// fmt.Print(l, ", ")
		r := r2(line)
		// fmt.Println(r, ", ")
		s.ans2 += r
		s.ans2 += 10 * l
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
