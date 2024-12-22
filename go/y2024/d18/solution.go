package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
)

type solution struct {
	ans int
	utils.Grid[bool]
}

type pt struct {
	utils.Pt
	step int
}

func (s *solution) run1() {
	// for _, row := range s.Array {
	// 	for _, v := range row {
	// 		if v {
	// 			fmt.Print("#")
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Println("")
	// }
	queue := []pt{{Pt: utils.Pt{C: 0, R: 0}, step: 0}}
	for len(queue) > 0 {
		n := len(queue)
		for n > 0 {
			n--
			l := queue[0]
			// fmt.Println(l)
			queue = queue[1:]
			if l.C == s.NCol-1 && l.R == s.NRow-1 {
				s.ans = l.step
				return
			}
			// if s.Get(l.Pt) {
			// 	continue
			// }
			for _, dir := range utils.Dir4 {
				nextP := l.PMove(dir)
				if !s.IsInside(nextP) || s.Get(nextP) {
					continue
				}
				s.Set(nextP, true)
				queue = append(queue, pt{Pt: nextP, step: l.step + 1})
			}
		}
	}
}

func (s *solution) run2() {
}

func (s *solution) res() int {
	return s.ans
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}
	nrow, ncol := 71, 71
	grid := make([][]bool, nrow)
	for i := range nrow {
		grid[i] = make([]bool, ncol)
	}
	for i, line := range lines {
		ints := utils.IntsFromString(line)
		grid[ints[1]][ints[0]] = true
		if i == 1023 {
			break
		}
	}

	return &solution{
		ans: 0,
		Grid: utils.Grid[bool]{
			NRow:  nrow,
			NCol:  ncol,
			Array: grid,
		},
	}
}

func part1(r io.Reader) int {
	s := buildSolution(r)
	s.run1()
	return s.res()
}

func part2(r io.Reader) int {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}
	nrow, ncol := 71, 71
	for checkedLines, res := range lines {
		if checkedLines <= 1023 {
			continue
		}

		grid := make([][]bool, nrow)
		for i := range nrow {
			grid[i] = make([]bool, ncol)
		}

		for i, line := range lines {
			ints := utils.IntsFromString(line)
			grid[ints[1]][ints[0]] = true
			if i == checkedLines {
				break
			}
		}

		s := &solution{
			ans: 0,
			Grid: utils.Grid[bool]{
				NRow:  nrow,
				NCol:  ncol,
				Array: grid,
			},
		}
		s.run1()
		if s.res() == 0 {
			fmt.Println(checkedLines, res)
			break
		}
	}
	return 0
}

func main() {
	Input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("fail open input.txt %v", err)
	}
	defer Input.Close()
	fmt.Println("p1 res 🙆-> ", part1(Input))
	fmt.Println("p2 res 🙆-> ", part2(Input))
}
