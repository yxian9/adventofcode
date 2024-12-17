package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
)

func Part1(r io.Reader) int {
	solution := buildSolution(r)
	solution.run()

	return solution.res()
}

func buildSolution(r io.Reader) solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v", err)
	}
	ante := antennas{}
	for i, line := range lines {
		for j, r := range line {
			if r == '.' {
				continue
			}
			ante[r] = append(ante[r], utils.Pt{X: i, Y: j})
		}
	}

	return solution{
		nrow:      len(lines),
		ncol:      len(lines[0]),
		antennas:  ante,
		antinodes: antinodes{},
	}
}

type (
	antennas  map[rune][]utils.Pt
	antinodes map[utils.Pt]struct{}
	solution  struct {
		nrow, ncol int
		antennas   antennas
		antinodes  antinodes
	}
)

func (s solution) res() int {
	return len(s.antinodes)
}

func (s solution) isInside(p utils.Pt) bool {
	return p.X >= 0 && p.X < s.nrow && p.Y >= 0 && p.Y < s.ncol
}

func (s solution) run() {
	for _, anteSlice := range s.antennas {
		for i := range anteSlice {
			for j := i + 1; j < len(anteSlice); j++ {

				ptSlice := [2]utils.Pt{
					anteSlice[i], anteSlice[j],
				}

				dx, dy := ptSlice[0].Dist(ptSlice[1])

				for i, pt := range ptSlice {
					npt := pt.Move(dx, dy)
					if i == 1 {
						npt = pt.Move(-dx, -dy)
					}
					if s.isInside(npt) {
						s.antinodes[npt] = struct{}{}
					}
				}
			}
		}
	}
}

func f(a int) int {
	if a == 0 {
		return 1
	}
	return -1
}

func (s solution) run2() {
	for _, antes := range s.antennas {
		for i := range antes {
			for j := i + 1; j < len(antes); j++ {
				// for each ante pairs
				atPair := [2]utils.Pt{
					antes[i], antes[j],
				}

				dx, dy := atPair[0].Dist(atPair[1])

				for i, at := range atPair {
					s.antinodes[at] = struct{}{}
					// cx, cy := dx, dy
					// if i == 1 {
					// 	cx, cy = -dx, -dy
					// }

					step := 1
					for {
						npt := at.Move(dx*step*f(i), dy*step*f(i))
						if s.isInside(npt) {
							s.antinodes[npt] = struct{}{}
							step++
						} else {
							break
						}
					}
				}
			}
		}
	}
}

func Part2(r io.Reader) int {
	solution := buildSolution(r)
	solution.run2()

	return solution.res()
}

func main() {
	arg := os.Args[1]
	fmt.Println("Running part", arg)
	switch arg {
	case "1":
		res := Part1(os.Stdin)
		fmt.Println("p1 res 🙆-> ", res)
	case "2":
		res := Part2(os.Stdin)
		fmt.Println("p2 res 🙆-> ", res)
	}
}
