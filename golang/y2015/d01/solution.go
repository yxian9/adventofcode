package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type solution struct {
	input []byte
	ans   int
}

func (s *solution) run1() {
	for _, v := range s.input {
		if v == byte(')') {
			s.ans--
		}
		if v == byte('(') {
			s.ans++
		}
	}
}

func (s *solution) run2() {
	for i, v := range s.input {
		if v == byte(')') {
			s.ans--
		}
		if v == byte('(') {
			s.ans++
		}
		if s.ans == -1 {
			s.ans = i + 1
			break
		}
	}
}

func (s *solution) res() int {
	return s.ans
}

func buildSolution(r io.Reader) *solution {
	lines, err := io.ReadAll(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}

	return &solution{
		input: lines,
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
	arg := os.Args[1]
	fmt.Println("Running part", arg)
	switch arg {
	case "1":
		fmt.Println("p1 res 🙆-> ", part1(os.Stdin))
	case "2":
		fmt.Println("p2 res 🙆-> ", part2(os.Stdin))
	}
}
