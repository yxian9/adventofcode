package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
)

type machine struct {
	a, b, target utils.Pt
}

type solution struct {
	ans      int
	limit    int
	machines []machine
}

func (s *solution) found(an, bn int, a, b, target utils.Pt) bool {
	return (an*a.X+bn*b.X) == target.X && (an*a.Y+bn*b.Y) == target.Y
}

func (s *solution) lowest(a, b, target utils.Pt) {
	lowest, alreadyFound := 0, false

	for aNum := range s.limit {
		for bNum := range s.limit {
			if s.found(aNum, bNum, a, b, target) {
				cur := 3*aNum + bNum
				if !alreadyFound {
					lowest = cur
				} else {
					lowest = min(lowest, cur)
				}
				alreadyFound = true
			}
		}
	}
	s.ans += lowest
}

func (s *solution) lowest2(a, b, target utils.Pt) {
	ax, ay, bx, by, px, py := a.X, a.Y, b.X, b.Y, target.X, target.Y
	det := ax*by - ay*bx
	fmt.Println(det, "det")

	if det == 0 {
		return
	}

	aPress := (px*by - py*bx) / det
	bPress := (ax*py - ay*px) / det

	if ax*aPress+bx*bPress == px && ay*aPress+by*bPress == py && aPress >= 0 && bPress >= 0 {
		cost := aPress*3 + bPress
		s.ans += cost
	}
}

func (s *solution) run1() {
	for _, machine := range s.machines {
		s.lowest(machine.a, machine.b, machine.target)
	}
}

func (s *solution) run2() {
	for _, machine := range s.machines {
		s.lowest2(machine.a, machine.b, machine.target)
	}
}

func (s *solution) res() int {
	return s.ans
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}
	machines := []machine{}
	ps := [3]utils.Pt{}

	for i, line := range lines {
		// fmt.Println(line)
		if i%4 == 3 {
			machines = append(machines, machine{
				ps[0], ps[1], ps[2],
			})
			continue
		}
		ints := utils.IntsFromString(line)
		ps[i%4].X = ints[0]
		ps[i%4].Y = ints[1]
	}

	machines = append(machines, machine{
		ps[0], ps[1], ps[2],
	})
	fmt.Println(len(machines), "machine lnes")

	return &solution{
		machines: machines,
		limit:    100,
		ans:      0,
	}
}

func buildSolution2(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}
	machines := []machine{}
	ps := [3]utils.Pt{}

	fmt.Println("part2", machines)
	fmt.Println(lines)
	for i, line := range lines {
		fmt.Println(line)
		if i%4 == 3 {
			machines = append(machines, machine{
				ps[0], ps[1], ps[2],
			})
			continue
		}
		ints := utils.IntsFromString(line)
		if i%4 == 2 {
			ps[i%4].X = ints[0] + 10000000000000
			ps[i%4].Y = ints[1] + 10000000000000
			continue
		}
		ps[i%4].X = ints[0]
		ps[i%4].Y = ints[1]
	}

	machines = append(machines, machine{
		ps[0], ps[1], ps[2],
	})
	fmt.Println(len(machines), "machine lnes")

	return &solution{
		machines: machines,
		limit:    100,
		ans:      0,
	}
}

func part1(r io.Reader) int {
	s := buildSolution(r)
	s.run1()
	return s.res()
}

func part2(r io.Reader) int {
	s := buildSolution2(r)
	s.run2()
	return s.res()
}

func main() {
	// fmt.Println("p1 res 🙆-> ", part1(os.Stdin))
	fmt.Println("p2 res 🙆-> ", part2(os.Stdin))
}
