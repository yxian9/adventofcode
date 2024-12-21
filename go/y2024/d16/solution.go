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
	minScore    map[utils.Pt]int
	posDirScore map[[2]utils.Pt]int
	paths       map[utils.Pt]bool
	start, end  utils.Pt
}
type bp struct {
	pos, dir utils.Pt
	score    int
	path     []utils.Pt
}

func (s *solution) dfs(p, angle utils.Pt, score int) {
	if s.PRune(p) == '#' || score > s.ans || s.Seen[p] {
		return
	}
	if s.PRune(p) == 'E' {
		if score == s.ans {
			fmt.Println("found", len(s.paths), "score:", score, "position", p)
			for k, v := range s.Seen {
				if v {
					s.paths[k] = true
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
		s.dfs(nextP, dir, nextScore)
	}
	s.Seen[p] = false
}

func (s *solution) posdir(p, d utils.Pt) (a [2]utils.Pt) {
	a[0], a[1] = p, d
	return a
}

func (s *solution) bfs(p, dir utils.Pt) {
	queue := []bp{{p, dir, 0, nil}}
	s.minScore[p] = 0

	for len(queue) > 0 {
		n := len(queue)

		for n > 0 {

			l := queue[0]
			queue = queue[1:]

			if s.PRune(l.pos) == 'E' {
				fmt.Println("first part E", "cur ans:", s.ans, "cur socre:", l.score)
				// if s.ans == 0 {
				// 	s.ans = l.score
				// } else {
				// 	s.ans = min(s.ans, l.score)
				// }
				if s.ans == 0 || l.score < s.ans {
					s.ans = l.score
				}
			}

			for _, dir := range utils.Dir4 {

				nextP := l.pos.PMove(dir)
				if s.PRune(nextP) == '#' {
					continue
				}
				turned := 0
				if l.dir != dir {
					turned = 1
				}
				next_score := l.score + 1 + 1000*turned

				score, ok := s.minScore[nextP]

				if !ok || next_score < score {
					s.minScore[nextP] = next_score
					queue = append(queue, bp{nextP, dir, next_score, nil})
				}

			}

			n--
		}
	}
}

func (s *solution) bfs2(p, dir utils.Pt) {
	queue := []bp{{p, dir, 0, nil}}
	for _, v := range utils.Dir4 {
		cur := s.posdir(p, v)
		s.posDirScore[cur] = 0
	}

	for len(queue) > 0 {

		n := len(queue)
		for n > 0 {
			l := queue[0]
			queue = queue[1:]
			n--
			if l.score > s.ans {
				continue
			}
			if s.PRune(l.pos) == 'E' && l.score == s.ans {

				for _, v := range l.path {
					s.paths[v] = true
				}
				// fmt.Println("found")
				continue
			}

			for _, dir := range utils.Dir4 {
				nextP := l.pos.PMove(dir)
				if s.PRune(nextP) == '#' {
					continue
				}
				nextPosDir := s.posdir(nextP, dir)
				turned := 0
				if l.dir != dir {
					turned = 1
				}
				next_score := l.score + 1 + 1000*turned
				score, ok := s.posDirScore[nextPosDir]

				if !ok || next_score <= score {
					s.posDirScore[nextPosDir] = next_score
					newPath := append([]utils.Pt{}, l.path...)
					newPath = append(newPath, l.pos)
					queue = append(queue, bp{nextP, dir, next_score, newPath})
					// queue = append(queue, bp{nextP, dir, next_score, append(l.path, l.pos)}) // this will modify the upderline array
					// if array get reallocated, when next dir use l.path, the l.path is gone is gone
				}
			}
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
	s.bfs2(s.start, utils.Dir4[3])
}

func (s *solution) res() int {
	return s.ans
}

func (s *solution) res2() int {
	return len(s.paths) + 1
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}

	return &solution{
		Seen:        make(utils.Seen),
		start:       utils.Pt{C: 1, R: len(lines) - 2},
		end:         utils.Pt{C: len(lines[0]) - 2, R: 1},
		minScore:    map[utils.Pt]int{},
		posDirScore: map[[2]utils.Pt]int{},
		paths:       map[utils.Pt]bool{},
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
