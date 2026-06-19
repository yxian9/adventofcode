package main

import (
	"adventofcode/h"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type solution struct {
	input                  h.StringGrid
	ans1, ans2, nrow, ncol int
	gearsM                 map[h.Pt][]int
}

func buildSolution(r io.Reader) *solution {
	lines, _ := h.LinesFromReader(r)
	return &solution{
		input:  h.StringGrid{Array: lines},
		nrow:   len(lines),
		ncol:   len(lines[0]),
		ans1:   0,
		ans2:   0,
		gearsM: make(map[h.Pt][]int),
	}
}

func isSymb(b byte) bool {
	if h.IsdigitBool(b) {
		return false
	}
	return b != '.'
}

func (s *solution) check_nearby(p h.Pt) int {
	for _, dir := range h.Dir8 {
		nxp := p.PMove(dir)
		if !s.input.IsInside(nxp) {
			continue
		}
		b := s.input.GetByte(nxp)
		if isSymb(b) {
			return 1
		}
	}
	return 0
}
func (s *solution) check_gear(p h.Pt) (gears []h.Pt) {
	for _, dir := range h.Dir8 {
		nxp := p.PMove(dir)
		if !s.input.IsInside(nxp) {
			continue
		}
		b := s.input.GetByte(nxp)
		if b == '*' {
			gears = append(gears, nxp)
		}
	}
	return
}

func (s *solution) traves() {
	for i := 0; i < s.nrow; i++ {
		for j := 0; j < s.ncol; j++ {

			pt := h.Pt{R: i, C: j}
			b := s.input.GetByte(pt)

			if h.IsdigitBool(b) {
				// start scanning
				num, valid := 0, 0

				// build num
				for h.IsdigitBool(b) {
					cur_d, _ := h.Isdigit(b)
					num = num*10 + cur_d
					valid += s.check_nearby(pt)
					j++
					pt = h.Pt{R: i, C: j}
					if !s.input.IsInside(pt) {
						break
					}
					b = s.input.GetByte(pt)
				}
				if valid > 0 {
					s.ans1 += num
				}
			}
		}
	}
}
func (s *solution) run1() {
	s.traves()
}

func (s *solution) traves2() {
	for i := 0; i < s.nrow; i++ {
		for j := 0; j < s.ncol; j++ {

			pt := h.Pt{R: i, C: j}
			b := s.input.GetByte(pt)

			if h.IsdigitBool(b) {
				// start scanning
				var (
					num   int
					gears = map[h.Pt]bool{}
				)

				// build num
				for h.IsdigitBool(b) {
					cur_d, _ := h.Isdigit(b)
					num = num*10 + cur_d
					for _, gear := range s.check_gear(pt) {
						gears[gear] = true
					}
					j++
					pt = h.Pt{R: i, C: j}
					if !s.input.IsInside(pt) {
						break
					}
					b = s.input.GetByte(pt)
				}
				for pt := range gears {
					s.gearsM[pt] = append(s.gearsM[pt], num)
				}
			}
		}
	}
}
func (s *solution) run2() {
	s.traves2()
	for _, ints := range s.gearsM {
		if len(ints) == 2 {
			s.ans2 += ints[0] * ints[1]
		}
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
