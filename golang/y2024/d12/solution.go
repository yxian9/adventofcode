package main

import (
	"adventofcode/golang/h"
	"fmt"
	"io"
	"log"
	"os"
)

type (
	headPt   h.Pt
	solution struct {
		ans int
		// lands map[headPt]map[utils.Pt]rune
		lands map[headPt][]h.Pt
		seen  map[h.Pt]bool
		h.StringGrid
	}
)

func (s *solution) run1() {
	for i, line := range s.Array {
		for j, flower := range line {
			cur := h.Pt{C: i, R: j}
			if s.seen[cur] {
				continue
			}
			s.dfs1(cur, flower, headPt(cur))
		}
	}
}

func (s *solution) res() int {
	for header, land := range s.lands {
		area, perimeter := len(land), 0
		flower := s.GetRune(h.Pt(header))
		for _, pt := range land {
			for _, dir := range h.Dir4 {
				nexPt := pt.Move(dir.C, dir.R)
				if !s.IsInside(nexPt) || s.GetRune(nexPt) != flower {
					perimeter++
				}
			}
		}
		s.ans += area * perimeter
	}
	return s.ans
}

func (s *solution) res2() int {
	for header, land := range s.lands {
		area, sides := len(land), 0
		flower := s.GetRune(h.Pt(header))
		for _, pt := range land {
			boolSlice := make([]bool, 4)
			// check whether neigher direc with same bool
			for i, dir := range h.Dir4 {
				nexPt := pt.Move(dir.C, dir.R)
				if !s.IsInside(nexPt) || s.GetRune(nexPt) != flower {
					boolSlice[i] = true
				}
			}
			for i, v := range boolSlice {
				neighb := (i + 1) % 4
				// out conner, both neighb directioin can not go
				if v && boolSlice[neighb] {
					sides++
				}
				// for insider conner, both can go
				if !v && !boolSlice[neighb] {
					pt1, pt2 := h.Dir4[i], h.Dir4[neighb]
					anglePt := pt.Move(pt1.C+pt2.C, pt1.R+pt2.R)
					if s.GetRune(anglePt) != flower {
						sides++
					}
				}
			}
		}
		// fmt.Printf("flower %c area %d slides %d \n", flower, area, sides)
		s.ans += area * sides
	}
	return s.ans
}

func (s *solution) dfs1(curP h.Pt, flower rune, header headPt) {
	if !s.IsInside(curP) || s.GetRune(curP) != flower || s.seen[curP] {
		return
	}
	s.seen[curP] = true
	// if s.lands[header] == nil {
	// 	s.lands[header] = map[utils.Pt]rune{}
	// }

	s.lands[header] = append(s.lands[header], curP)
	for _, dir := range h.Dir4 {
		nextP := curP.Move(dir.C, dir.R)
		s.dfs1(nextP, flower, header)
	}
}

func (s *solution) run2() {
}

func buildSolution(r io.Reader) *solution {
	lines, err := h.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}
	nrow, ncol := len(lines), len(lines[0])

	return &solution{
		ans:   0,
		seen:  map[h.Pt]bool{},
		lands: map[headPt][]h.Pt{},
		StringGrid: h.StringGrid{
			NRow:  nrow,
			NCol:  ncol,
			Array: lines,
		},
	}
}

func part1(r io.Reader) int {
	s := buildSolution(r)
	s.run1()
	return s.res()
}

func part2(r io.Reader) int {
	s := buildSolution(r)
	s.run1()
	return s.res2()
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
