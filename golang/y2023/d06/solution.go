package main

import (
	"adventofcode/golang/h"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type solution struct {
	input       []string
	ans1, ans2  int
	time, dis   []int
	time2, dis2 []int
}

func parse(s string, trim bool) []int {
	_, remain, _ := strings.Cut(s, ":")
	if trim {
		remain = strings.ReplaceAll(remain, " ", "")
	}
	return h.IntsFromString(remain)
}
func buildSolution(r io.Reader) *solution {
	lines, _ := h.LinesFromReader(r)
	return &solution{
		input: lines,
		ans1:  0,
		ans2:  0,
		time:  parse(lines[0], false),
		dis:   parse(lines[1], false),
		time2: parse(lines[0], true),
		dis2:  parse(lines[1], true),
	}
}

func allPossib(t, dis int) (res int) {
	for i := range t {
		if i*(t-i) > dis {
			res += 1
		}
	}
	return res
}

func (s *solution) run1() {
	var res = []int{}
	for i, v := range s.time {
		res = append(res, allPossib(v, s.dis[i]))
	}
	s.ans1 = 1
	for _, v := range res {
		s.ans1 *= v
	}
}

func (s *solution) run2() {
	s.ans2 = allPossib(s.time2[0], s.dis2[0])
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

	_, filename, _, _ := runtime.Caller(0)
	Input, err := os.Open(filepath.Join(filepath.Dir(filename), "input.txt"))
	if err != nil {
		log.Fatalf("fail open input.txt %v", err)
	}
	start := time.Now()
	result := part1(Input)
	elapsed := time.Since(start)
	fmt.Printf("p1 res 🙆-> %d (Time taken: %s)\n", result, elapsed)
	Input.Close()
	Input, err = os.Open(filepath.Join(filepath.Dir(filename), "input.txt"))
	start = time.Now()
	result = part2(Input)
	elapsed = time.Since(start)
	fmt.Printf("p2 res 🙆-> %d (Time taken: %s)\n", result, elapsed)
	Input.Close()
}
