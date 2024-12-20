package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
)

type solution struct {
	ans, step int
	utils.StringGrid
	utils.Seen
	count      map[utils.Pt]int
	space      map[utils.Pt]bool
	start, end utils.Pt
}
type bp struct {
	p, d  utils.Pt
	score int
}

func (s *solution) dfs(p, angle utils.Pt, score int) {
	if s.PRune(p) == '#' || score > s.ans || s.Seen[p] {
		return
	}
	if s.PRune(p) == 'E' {
		if score == s.ans {
			fmt.Println("found", len(s.space), "score:", score, "position", p)
			for k, v := range s.Seen {
				if v {
					s.space[k] = true
				}
			}
		}
		return
	}

	s.Seen[p] = true
	for _, dir := range utils.Dir4 {
		nextP := p.PMove(dir)
		turned := 0
		if angle != dir {
			turned = 1
		}
		nextScore := score + 1 + 1000*turned
		// if s.count[nextP] == nextScore {
		s.dfs(nextP, dir, nextScore)
		// }
	}
	s.Seen[p] = false
}

func (s *solution) bfs(p, angle utils.Pt) {
	queue := []bp{{p, angle, 0}}
	s.count[p] = 0
	for len(queue) > 0 {
		n := len(queue)

		for n > 0 {

			l := queue[0]
			queue = queue[1:]

			if s.PRune(l.p) == 'E' {
				// fmt.Println("bfs E", "cur ans:", s.ans, "cur socre:", l.score)
				if s.ans == 0 {
					s.ans = l.score
				} else {
					s.ans = min(s.ans, l.score)
				}
			}

			for _, dir := range utils.Dir4 {

				nextP := l.p.PMove(dir)
				if !s.IsInside(nextP) || s.PRune(nextP) == '#' {
					continue
				}
				turned := 0
				if l.d != dir {
					turned = 1
				}
				next_score := l.score + 1 + 1000*turned

				preScore, ok := s.count[nextP]

				if !ok || next_score <= preScore {
					s.count[nextP] = next_score
					queue = append(queue, bp{nextP, dir, next_score})
				}

			}

			n--
		}
	}
}

func (s *solution) bfs2(p, angle utils.Pt) {
	queue := []bp{{p, angle, s.ans}}
	// fmt.Println(s.count[p], p)
	for len(queue) > 0 {
		n := len(queue)

		for n > 0 {
			l := queue[0]
			queue = queue[1:]
			s.space[l.p] = true
			nextTarget := l.score
			for _, dir := range utils.Dir4 {
				nextP := l.p.PMove(dir)
				if s.space[nextP] {
					continue
				}

				if nextScore, ok := s.count[nextP]; ok {
					nextTarget = min(nextTarget, nextScore)
				}
			}
			for _, dir := range utils.Dir4 {
				nextP := l.p.PMove(dir)
				if s.space[nextP] || s.count[nextP] == 0 {
					continue
				}
				if s.count[nextP] == nextTarget || s.count[nextP] == nextTarget+1000 {
					// fmt.Println("nexttarget:", nextTarget, "enqueued:", s.count[nextP])
					queue = append(queue, bp{nextP, dir, nextTarget})
				}
			}

			n--
		}
	}
}

func (s *solution) run1() {
	s.bfs(s.start, utils.Dir4[3])
	// debug the count map

	// for r, line := range s.Array {
	// 	for c, ru := range line {
	// 		p := utils.Pt{C: c, R: r}
	// 		if score, ok := s.count[p]; ok {
	// 			fmt.Printf("%6d", score)
	// 		} else {
	// 			fmt.Printf("%6c", ru)
	// 		}
	// 	}
	// 	fmt.Println("")
	// }
}

func (s *solution) run2() {
	s.bfs(s.start, utils.Dir4[3])
	s.bfs2(s.end, utils.Dir4[3])
	// s.dfs(s.start, utils.Dir4[3], 0)
}

func (s *solution) res() int {
	return s.ans
}

func (s *solution) res2() int {
	return len(s.space) + 1
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}

	return &solution{
		Seen:  make(utils.Seen),
		start: utils.Pt{C: 1, R: len(lines) - 2},
		end:   utils.Pt{C: len(lines[0]) - 2, R: 1},
		count: map[utils.Pt]int{},
		space: map[utils.Pt]bool{},
		StringGrid: utils.StringGrid{
			NRow:  len(lines),
			NCol:  len(lines[0]),
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
