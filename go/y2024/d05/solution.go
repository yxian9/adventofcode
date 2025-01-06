package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"time"
)

func (s *solution) run1() {
	for _, manual := range s.manuals {
		if s.isSafe(manual) {
			s.ans += manual[len(manual)/2]
		}
	}
}

func (s *solution) isSafe(manual []int) bool {
	parents := map[int]bool{}
	for _, v := range manual {
		for _, child := range s.graph[v] {
			if parents[child] {
				return false
			}
		}
		parents[v] = true
	}
	return true
}

func (s *solution) builInDegree(graph map[int][]int) (inDegree map[int]int) {
	inDegree = map[int]int{}
	for _, childen := range graph {
		for _, child := range childen {
			inDegree[child]++
		}
	}
	return inDegree
}

func (s *solution) buildGraph(manual []int) (graph map[int][]int) {
	graph = map[int][]int{}
	for _, parent := range manual {
		var dep []int
		for _, child := range s.graph[parent] {
			if slices.Contains(manual, child) {
				dep = append(dep, child)
			}
		}
		graph[parent] = dep
	}
	return graph
}

func (s *solution) topSort(manual []int) (sortedM []int) {
	graph := s.buildGraph(manual)
	inDegree := s.builInDegree(graph)
	queue := []int{}
	for _, v := range manual {
		if inDegree[v] == 0 {
			queue = append(queue, v)
		}
	}
	for len(queue) > 0 {
		parent := queue[0]
		sortedM = append(sortedM, parent)
		queue = queue[1:]
		for _, child := range graph[parent] {
			inDegree[child]--
			if inDegree[child] == 0 {
				queue = append(queue, child)
			}
		}
	}
	return sortedM
}

func (s *solution) run2() {
	for _, manual := range s.manuals {
		if !s.isSafe(manual) {
			// fmt.Println(manual)
			sorted_manual := s.topSort(manual)
			s.ans += sorted_manual[len(sorted_manual)/2]
		}
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
	first := true
	graph := make(map[int][]int, 0)
	manuals := [][]int{}
	for _, line := range lines {
		if len(line) == 0 {
			first = false
			continue
		}
		ints := utils.IntsFromString(line)
		if first {
			graph[ints[0]] = append(graph[ints[0]], ints[1])
		} else {
			manuals = append(manuals, ints)
		}
	}

	return &solution{
		graph:   graph,
		manuals: manuals,
		ans:     0,
	}
}

type solution struct {
	manuals [][]int
	graph   map[int][]int
	ans     int
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
