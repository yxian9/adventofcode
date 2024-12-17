package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
)

type robot struct {
	utils.Pt
	vx, vy int
}

type solution struct {
	ans    int
	robots []robot
}

const (
	ncol, nrow = 101, 103
	seconds    = 100
)

func (s *solution) run1() {
	for i := range s.robots {
		s.robots[i].X = (s.robots[i].X + s.robots[i].vx*seconds) % ncol
		s.robots[i].Y = (s.robots[i].Y + s.robots[i].vy*seconds) % nrow
		if s.robots[i].X < 0 {
			s.robots[i].X += ncol
		}
		if s.robots[i].Y < 0 {
			s.robots[i].Y += nrow
		}
	}
}

func (s *solution) treeCheck(sec int) bool {
	pixels := make([][]bool, nrow)
	for row := range pixels {
		pixels[row] = make([]bool, ncol)
	}

	for _, r := range s.robots {
		pixels[r.Y][r.X] = true
	}

	// if sec%1000 == 0 {
	// 	fmt.Println("###########################################")
	// 	for i := range nrow {
	// 		for j := range ncol {
	// 			if pixels[i][j] {
	// 				fmt.Printf("#")
	// 			} else {
	// 				fmt.Printf(" ")
	// 			}
	// 		}
	// 		fmt.Println("")
	// 	}
	// }

	for row := range pixels {
		consecutivePixels := 0
		for col := range pixels[row] {
			if pixels[row][col] {
				consecutivePixels++
			} else {
				consecutivePixels = 0
			}
			if consecutivePixels == 20 {
				fmt.Println("############################################")
				for i := range nrow {
					for j := range ncol {
						if pixels[i][j] {
							fmt.Printf("#")
						} else {
							fmt.Printf(" ")
						}
					}
					fmt.Println("")
				}
				return true
			}
		}
	}
	return false
}

func (s *solution) run2() {
	sec := 0
	for {
		for i := range s.robots {
			s.robots[i].X = (s.robots[i].X + s.robots[i].vx) % ncol
			s.robots[i].Y = (s.robots[i].Y + s.robots[i].vy) % nrow
			if s.robots[i].X < 0 {
				s.robots[i].X += ncol
			}
			if s.robots[i].Y < 0 {
				s.robots[i].Y += nrow
			}
		}
		sec++
		if s.treeCheck(sec) {
			s.ans = sec
			return
		}
	}
}

func (s *solution) res() int {
	quadrants := [4]int{}
	for _, r := range s.robots {
		if r.X < ncol/2 && r.Y < nrow/2 {
			quadrants[0]++
		} else if r.X < ncol/2 && r.Y > nrow/2 {
			quadrants[1]++
		} else if r.X > ncol/2 && r.Y > nrow/2 {
			quadrants[2]++
		} else if r.X > ncol/2 && r.Y < nrow/2 {
			quadrants[3]++
		}
	}
	s.ans = 1
	for i := range 4 {
		s.ans *= quadrants[i]
	}
	// s.ans = quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
	return s.ans
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}
	robots := []robot{}

	for _, line := range lines {
		ints := utils.IntsFromString(line)
		rb := robot{
			utils.Pt{X: ints[0], Y: ints[1]},
			ints[2], ints[3],
		}
		robots = append(robots, rb)
	}
	return &solution{
		ans:    0,
		robots: robots,
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
	return s.ans
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
