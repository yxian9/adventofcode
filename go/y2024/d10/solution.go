package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
)

type solution struct {
	nrow, ncol  int
	grid        []string
	seen        map[utils.Pt]bool
	ans         int
	trailHeader map[string]bool
}

func (s *solution) isInside(pt utils.Pt) bool {
	return pt.X >= 0 && pt.X < s.nrow && pt.Y >= 0 && pt.Y < s.ncol
}

func (s *solution) found(sp, cp utils.Pt) {
	id := fmt.Sprintf("%d,%d,%d,%d", sp.X, sp.Y, cp.X, cp.Y)
	s.trailHeader[id] = true
}

func (s *solution) pInt(pt utils.Pt) int {
	return int(s.grid[pt.X][pt.Y] - '0')
}

func (s *solution) dfs1(startP, curP utils.Pt, target int) {
	if !s.isInside(curP) || s.seen[curP] || s.pInt(curP) != target {
		return
	}
	if target == 9 {
		s.found(startP, curP)
		return
	}
	s.seen[curP] = true
	for _, dir := range utils.Dir4 {
		s.dfs1(startP, curP.Move(dir.X, dir.Y), target+1)
	}
	s.seen[curP] = false
}

func (s *solution) run1() {
	for i, line := range s.grid {
		for j, r := range line {
			if r == '0' {
				cp := utils.Pt{X: i, Y: j}
				s.dfs1(cp, cp, 0)
			}
		}
	}
}

func (s *solution) dfs2(startP, curP utils.Pt, target int) {
	if !s.isInside(curP) || s.seen[curP] || s.pInt(curP) != target {
		return
	}
	if s.pInt(curP) == 9 {
		s.ans++
		return
	}
	for _, dir := range utils.Dir4 {
		nxP := curP.Move(dir.X, dir.Y)
		s.seen[curP] = true
		s.dfs2(startP, nxP, target+1)
		s.seen[curP] = false
	}
}

func (s *solution) run2() {
	for i, line := range s.grid {
		for j, r := range line {
			if r == '0' {
				cp := utils.Pt{X: i, Y: j}
				s.dfs2(cp, cp, 0)
			}
		}
	}
}

func (s *solution) res() int {
	return len(s.trailHeader)
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	nr, nc := len(lines), len(lines[0])
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}
	seen := make([][]bool, nr)
	for i := range nr {
		seen[i] = make([]bool, nc)
	}

	return &solution{
		nrow:        nr,
		ncol:        nc,
		seen:        make(map[utils.Pt]bool),
		grid:        lines,
		trailHeader: make(map[string]bool),
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
	return s.ans
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
