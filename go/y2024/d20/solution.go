package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

const (
	W = '#'
	S = 'S'
	E = 'E'
	T = '.'
)

type solution struct {
	start utils.Pt
	utils.Seen
	utils.StringGrid
	ans    int
	saves  map[int]int
	orders map[utils.Pt]int
}

func (s *solution) makeOrder(p utils.Pt, step int) {
	// invalid
	if s.GetRune(p) == W || s.Seen[p] {
		return
	}
	s.orders[p] = step
	if s.GetRune(p) == E {
		return
	}

	s.Seen[p] = true
	for _, dir := range utils.Dir4 {
		nextP := p.PMove(dir)
		s.makeOrder(nextP, step+1)
	}
	s.Seen[p] = false
}

func (s *solution) findCheat(p utils.Pt) (cheatS utils.Pt, cheatEs []utils.Pt, steps []int, find bool) {
	for _, dir := range utils.Dir4 {
		nextP := p.Move(dir.R*2, dir.C*2)
		if s.IsInside(nextP) && s.GetRune(nextP) != W && !s.Seen[nextP] {
			cheatEs = append(cheatEs, nextP)
		}
	}
	if len(cheatEs) > 0 {
		return p, cheatEs, nil, true
	}
	return cheatS, nil, nil, false
}

func (s *solution) findCheat2(start utils.Pt) {
	queue := []utils.Pt{start}
	seen := map[utils.Pt]bool{}
	seen[start] = true
	step := 0

	for len(queue) > 0 {

		n := len(queue)
		for n > 0 {
			n--
			p := queue[0]
			queue = queue[1:]
			if s.GetRune(p) != W && p != start {
				distance := s.orders[p] - s.orders[start]
				if save := distance - step; save >= 50 {
					s.saves[save]++
				}
			}
			for _, dir := range utils.Dir4 {
				nextP := p.Move(dir.R, dir.C)
				if s.IsInside(nextP) && !seen[nextP] {
					seen[nextP] = true
					queue = append(queue, nextP)
				}
			}
		}
		step++
		if step > 20 {
			break
		}
	}
}

type findCheat func(utils.Pt) (utils.Pt, []utils.Pt, []int, bool)

func (s *solution) walk(p utils.Pt, findCheat findCheat) {
	if s.GetRune(p) == W || s.Seen[p] {
		return
	}
	if s.GetRune(p) == E {
		return
	}

	start, cheatEs, steps, find := findCheat(p)
	if find {
		// fmt.Printf("start %#v cheeat %v \n ", start, cheatEs)
		for i, cheatEnd := range cheatEs {
			diff := 2
			if steps != nil {
				diff = steps[i]
			}
			distant := s.orders[cheatEnd] - s.orders[start]
			s.saves[distant-diff]++
		}
	}

	s.Seen[p] = true
	for _, dir := range utils.Dir4 {
		nextP := p.PMove(dir)
		s.walk(nextP, findCheat)
	}
	s.Seen[p] = false
}

func (s *solution) walk2(p utils.Pt) {
	if s.GetRune(p) == W || s.Seen[p] {
		return
	}
	if s.GetRune(p) == E {
		return
	}

	s.findCheat2(p)

	s.Seen[p] = true
	for _, dir := range utils.Dir4 {
		nextP := p.PMove(dir)
		s.walk2(nextP)
	}
	s.Seen[p] = false
}

func (s *solution) run1() {
	s.makeOrder(s.start, 0)
	s.walk(s.start, s.findCheat)
}

func (s *solution) run2() {
	s.makeOrder(s.start, 0)
	s.walk2(s.start)
}

func (s *solution) res() int {
	for k, v := range s.saves {
		if k >= 100 {
			s.ans += v
		}
	}
	return s.ans
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}
	var start utils.Pt
	for r, row := range lines {
		for c, v := range row {
			if v == S {
				start = utils.Pt{C: c, R: r}
			}
		}
	}

	return &solution{
		start:  start,
		Seen:   utils.Seen{},
		saves:  map[int]int{},
		orders: map[utils.Pt]int{},
		StringGrid: utils.StringGrid{
			Array: lines,
		},
		ans: 0,
	}
}

func part1(r io.Reader) int {
	s := buildSolution(r)
	s.run1()
	// fmt.Printf("part1 %#v \n", s.saves)
	return s.res()
}

func part2(r io.Reader) int {
	s := buildSolution(r)
	s.run2()

	// fmt.Printf("part2 %#v \n", s.saves)
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
