package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type solution struct {
	nums []int
	ans  int
	memo map[string]int
}

func (s *solution) blink(stone int) (newStones []int) {
	if stone == 0 {
		newStones = append(newStones, 1)
	} else if len(strconv.Itoa(stone))%2 == 0 {
		numStr := strconv.Itoa(stone)
		half := len(numStr) / 2
		left, _ := strconv.Atoi(numStr[:half])
		right, _ := strconv.Atoi(numStr[half:])
		newStones = append(newStones, left, right)
	} else {
		newStones = append(newStones, stone*2024)
	}
	return newStones
}

func (s *solution) dfs(stone, level, target int) int {
	if level == target {
		return 1
	}
	id := fmt.Sprintf("%d,%d", stone, level)
	if v, ok := s.memo[id]; ok {
		return v
	}
	newStones := s.blink(stone)
	ans := 0
	for _, v := range newStones {
		ans += s.dfs(v, level+1, target)
	}
	s.memo[id] = ans
	return ans
}

func (s *solution) run1() {
	for _, v := range s.nums {
		s.ans += s.dfs(v, 0, 75)
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
	nums := utils.IntsFromString(lines[0])

	return &solution{
		nums: nums,
		memo: map[string]int{},
		ans:  0,
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
	return s.res()
}

func main() {
	arg := os.Args[1]
	fmt.Println("Running part", arg)
	switch arg {
	case "1":
		fmt.Println("p1 res 🙆-> ", part1(os.Stdin))
	case "2":
		fmt.Println("p2 res 🙆-> ", part2(os.Stdin))
	}
}
