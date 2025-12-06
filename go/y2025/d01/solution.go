package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

type solution struct {
	dial, len  int
	input      []string
	ans1, ans2 int
}

func parser(s string) (i, dir int) {
	i, _ = strconv.Atoi(s[1:])
	if s[0] == 'L' {
		dir = -1
	} else {
		dir = 1
	}
	return i, dir
}

func divmod(i, j int) (div, mod int) {
	div = i / j
	mod = i % j
	return
}

func (s *solution) move(rot string) {
	turn, dir := parser(rot)
	div, mod := divmod(turn, s.len)
	s.ans2 += div
	oldPos := s.dial
	s.dial += (mod * dir)
	switch {
	case s.dial == 0:
		s.ans2++
		// s.ans1++ // cannot single add at here. button condition can also add1
	case s.dial >= s.len:
		s.ans2++
		s.dial -= s.len
		// if s.dial == 0 {
		// 	s.ans1++
		// }
	case s.dial < 0:
		// If I'm starting from 0, I should not count the passage
		if oldPos != 0 {
			s.ans2++
		}
		s.dial += s.len
	}

	if s.dial == 0 {
		s.ans1++
	}
}

func (s *solution) run1() {
	for _, instru := range s.input {
		s.move(instru)
	}
}

func (s *solution) run2() {
	fmt.Println("Start Part2")
	for _, instru := range s.input {
		s.move(instru)
	}
}

func (s *solution) res1() int {
	return s.ans1
}

func (s *solution) res2() int {
	fmt.Println("res2", s.ans2)
	return s.ans2
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}

	return &solution{
		input: lines,
		dial:  50,
		len:   100,
	}
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
