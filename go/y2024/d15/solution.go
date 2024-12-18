package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
)

type solution struct {
	ans      int
	robot    utils.Pt
	instruct []rune
	utils.Grid[rune]
}

var Dirs = map[rune]utils.Pt{
	'v': {X: 0, Y: 1},
	'<': {X: -1, Y: 0},
	'^': {X: 0, Y: -1},
	'>': {X: 1, Y: 0},
}

func (s *solution) run1() {
	s.Array[s.robot.Y][s.robot.X] = '.'
	for _, r := range s.instruct {
		dir := Dirs[r]
		nexMove := s.robot.PMove(dir)
		nextR := s.Get(nexMove)
		if nextR == '#' {
			continue
		}
		if nextR == '.' {
			s.robot = nexMove
			continue
		}
		if nextEmptSpot, ok := s.Find(s.robot, dir, '.', '#'); ok {
			s.Swap(nextEmptSpot, nexMove)
			s.robot = nexMove
		}
	}
}

func (s *solution) run2() {
}

func (s *solution) res() int {
	for i, row := range s.Array {
		for j, r := range row {
			if r == 'O' {
				s.ans += i*100 + j
			}
		}
	}
	return s.ans
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}
	var robot utils.Pt
	nrow, ncol := 0, len(lines[0])
	firstPart := true

	grid, instruc := [][]rune{}, []rune{}

	for i, line := range lines {
		if len(line) == 0 {
			firstPart = false
			nrow = i
			continue
		}
		if firstPart {
			for j, r := range line {
				if r == '@' {
					robot = utils.Pt{X: i, Y: j}
				}
			}
			grid = append(grid, []rune(line))
		} else {
			instruc = append(instruc, []rune(line)...)
		}
	}

	return &solution{
		instruct: instruc,
		robot:    robot,
		Grid: utils.Grid[rune]{
			NRow:  nrow,
			NCol:  ncol,
			Array: grid,
		},
	}
}

func part1(r io.Reader) int {
	s := buildSolution(r)
	// fmt.Printf("%#v", s)
	s.run1()
	return s.res()
}

func part2(r io.Reader) int {
	s := buildSolution(r)
	s.run2()
	return s.res()
}

func main() {
	fmt.Println("start")
	Input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("fail open input.txt %v", err)
	}
	defer Input.Close()
	fmt.Println("p1 res 🙆-> ", part1(Input))
	// fmt.Println("p2 res 🙆-> ", part2(Input))
}
