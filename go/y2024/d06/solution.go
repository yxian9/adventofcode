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
	s.dfs(s.start, 0)
}

func (s *solution) dfs(cur utils.Pt, angle int) {
	if !s.IsInside(cur) {
		return
	}
	s.seen[cur] = true
	dir := utils.Dir4[angle]
	nextP := cur.PMove(dir)
	if !s.IsInside(nextP) {
		s.dfs(nextP, angle)
	} else if s.GetRune(nextP) == '#' {
		s.dfs(cur, (angle+1)%4)
	} else {
		s.dfs(nextP, angle)
	}
}

func (s *solution) original_dfs(start utils.Pt, angle int) {
	s.seen[start] = true // the logic make sure the start is inside
	dir := utils.Dir4[angle]
	nextP := start.PMove(dir)
	if !s.IsInside(nextP) {
		return
	}
	if s.GetRune(nextP) == '#' {
		s.dfs(start, (angle+1)%4)
	} else {
		s.dfs(nextP, angle)
	}
}

func (s *solution) dfs2(start utils.Pt, angle int, visit map[string]bool) {
	coordinate := fmt.Sprintf("%v : %d", start, angle)
	if visit[coordinate] {
		s.ans++
		return
	}
	visit[coordinate] = true
	dir := utils.Dir4[angle]
	nextP := start.PMove(dir)
	if !s.IsInside(nextP) {
		return
	}
	if s.GetRune(nextP) == '#' || s.block[nextP] {
		s.dfs2(start, (angle+1)%4, visit)
	} else {
		s.dfs2(nextP, angle, visit)
	}
}

func (s *solution) run2() {
	s.dfs(s.start, 0)
	for r, line := range s.Array {
		for c := range line {
			pt := utils.Pt{R: r, C: c}
			if s.seen[pt] {
				s.block[pt] = true
				s.dfs2(s.start, 0, map[string]bool{})
				s.block[pt] = false
			}
		}
	}
}

func (s *solution) res() int {
	return len(s.seen)
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}
	var start utils.Pt
	for r, line := range lines {
		for c, char := range line {
			if char == '^' {
				start = utils.Pt{C: c, R: r}
			}
		}
	}

	return &solution{
		StringGrid: utils.StringGrid{Array: lines},
		start:      start,
		seen:       map[utils.Pt]bool{},
		block:      map[utils.Pt]bool{},
	}
}

type solution struct {
	utils.StringGrid
	start utils.Pt
	seen  map[utils.Pt]bool
	block map[utils.Pt]bool
	ans   int
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
