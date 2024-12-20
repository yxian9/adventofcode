package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
)

type solution struct {
	ans      int
	robot    utils.Pt
	instruct []rune
	utils.Grid[rune]
}

var Dirs = map[rune]utils.Pt{
	'v': {C: 0, R: 1},
	'<': {C: -1, R: 0},
	'^': {C: 0, R: -1},
	'>': {C: 1, R: 0},
}

func (s *solution) run1() {
	s.Array[s.robot.R][s.robot.C] = '.'
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

func (s *solution) loopSwap(ptSlice []utils.Pt, dir utils.Pt) {
	for i := len(ptSlice) - 1; i >= 0; i-- {
		nextPt := ptSlice[i].PMove(dir)
		s.Swap(nextPt, ptSlice[i])
	}
}

func (s *solution) getPair(pt utils.Pt) (ptSlice []utils.Pt) {
	if s.Get(pt) == '[' {
		ptSlice = append(ptSlice, pt.Move(1, 0), pt)
		return ptSlice
	}
	ptSlice = append(ptSlice, pt.Move(-1, 0), pt)
	return ptSlice
}

func (s *solution) yfind(start utils.Pt, dir utils.Pt) (ptSlice []utils.Pt, find bool) {
	var queue []utils.Pt
	queue = append(queue, s.getPair(start)...)
	for len(queue) > 0 {
		n := len(queue)
		for n > 0 {
			l1 := queue[0]
			queue = queue[1:]
			ptSlice = append(ptSlice, l1)
			nextPt := l1.PMove(dir)
			switch s.Get(nextPt) {
			case '#':
				return nil, false
			case '[', ']':
				for _, v := range s.getPair(nextPt) {
					if !slices.Contains(queue, v) {
						queue = append(queue, v)
					}
				}
			case '.':

			}
			n--
		}
	}
	return ptSlice, true
}

func (s *solution) xfind(start utils.Pt, dir utils.Pt) (ptSlice []utils.Pt, find bool) {
	for {
		ptSlice = append(ptSlice, start)
		nextPt := start.PMove(dir)

		switch s.Get(nextPt) {
		case '#':
			return nil, false
		case '.':
			return ptSlice, true
		}

		start = nextPt
	}
}

func (s *solution) run2() {
	// fmt.Println("-----")
	// for _, line := range s.Array {
	// 	fmt.Println(string(line))
	// }

	for _, r := range s.instruct {
		dir := Dirs[r]
		nexMove := s.robot.PMove(dir)
		nextR := s.Get(nexMove)
		// time.Sleep(1 * time.Second)
		// fmt.Printf("cur %v, next %v dir: %c,find %c \n", s.robot, nexMove, r, nextR)
		// fmt.Println("-----")
		// for y, line := range s.Array {
		// 	for x, r := range line {
		// 		if x == s.robot.X && y == s.robot.Y {
		// 			r = '@'
		// 		}
		// 		fmt.Printf("%c", r)
		// 	}
		// 	fmt.Println("")
		// }
		if nextR == '#' {
			continue
		}
		if nextR == '.' {
			s.robot = nexMove
			continue
		}
		switch r {
		case '^', 'v':
			if ptslice, find := s.yfind(nexMove, dir); find {
				s.loopSwap(ptslice, dir)
				s.robot = nexMove
			}
		case '>', '<':
			if ptslice, find := s.xfind(nexMove, dir); find {
				s.loopSwap(ptslice, dir)
				s.robot = nexMove
			}
		}
	}
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

func (s *solution) res2() int {
	for i, row := range s.Array {
		for j, r := range row {
			if r == '[' {
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
					robot = utils.Pt{C: j, R: i}
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

func buildSolution2(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}
	var robot utils.Pt
	nrow, ncol := 0, 2*len(lines[0])
	firstPart := true

	grid, instruc := [][]rune{}, []rune{}

	for i, line := range lines {
		if len(line) == 0 {
			firstPart = false
			nrow = i
			continue
		}
		if firstPart {
			cur := []rune{}
			for j, r := range line {
				switch r {
				case '@':
					robot = utils.Pt{C: 2 * j, R: i}
					cur = append(cur, '.', '.')
				case '#':
					cur = append(cur, '#', '#')
				case 'O':
					cur = append(cur, '[', ']')
				case '.':
					cur = append(cur, '.', '.')
				}
			}

			grid = append(grid, cur)
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
	s := buildSolution2(r)
	s.run2()
	return s.res2()
}

func main() {
	fmt.Println("start")
	Input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("fail open input.txt %v", err)
	}
	defer Input.Close()
	fmt.Println("p1 res 🙆-> ", part1(Input))
	fmt.Println("p2 res 🙆-> ", part2(Input))
}
