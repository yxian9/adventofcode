package main

import (
	"adventofcode/golang/h"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
type cube struct {
	red, green, blue int
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

func parse(s string) (id int, cubes []cube) {
	g, rst, _ := strings.Cut(s, ":")
	fmt.Sscanf(g, "Game %d", &id)
	for set := range strings.SplitSeq(rst, ";") {
		for c := range strings.SplitSeq(set, ",") {
			var (
				i    int
				col  string
				cube cube
			)
			fmt.Sscanf(c, "%d %s", &i, &col)
			switch col {
			case "blue":
				cube.blue = i
			case "red":
				cube.red = i
			case "green":
				cube.green = i
			}
			cubes = append(cubes, cube)
		}
	}
	return
}
func valid(lst []cube) bool {
	for _, c := range lst {
		if c.blue > 14 || c.red > 12 || c.green > 13 {
			return false
		}
	}
	return true
}
func sum(lst []cube) int {
	var r, g, b int
	for _, c := range lst {
		r = max(c.red, r)
		g = max(c.green, g)
		b = max(c.blue, b)
	}
	return r * g * b

}

func (s *solution) run1() {
	for _, v := range s.input {
		id, cubes := parse(v)
		if valid(cubes) {
			s.ans1 += id
		}
	}
}

func (s *solution) run2() {
	for _, v := range s.input {
		_, cubes := parse(v)
		s.ans2 += sum(cubes)
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
