package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"sort"
	"strings"
	"time"
)

type solution struct {
	relations map[string][]string
	ans       int
	connet3   map[string]bool
}

func (s *solution) run1() {
	for a, a_childs := range s.relations {
		if len(a_childs) < 2 {
			continue
		}
		if a[0] != 't' {
			continue
		}
		for i, b := range a_childs {
			// if a <= b {
			// 	continue
			// }
			for j := i + 1; j < len(a_childs); j++ {
				c := a_childs[j] // both b and c need to be a's child
				if slices.Contains(s.relations[b], c) {
					keyslice := []string{a, b, c}
					sort.Strings(keyslice)
					key := fmt.Sprint(keyslice)
					// fmt.Println(key)
					s.connet3[key] = true
				}
			}
		}
	}

	var largest []string
	for a, v := range s.relations {
		cohort := s.dfs_expand([]string{a}, v)
		if len(cohort) > len(largest) {
			largest = cohort
		}
	}
	slices.Sort(largest)
	fmt.Println(strings.Join(largest, ","))
}

func (s *solution) dfs_expand(current, candidates []string) []string {
	// If we're unable to expand, return just what we have.
	if len(candidates) == 0 {
		return current
	}
	largest := current
	proposed := make([]string, len(current)+1)
	copy(proposed, current)
outer:
	for i, c := range candidates {
		for _, v := range current {
			if !slices.Contains(s.relations[v], c) {
				continue outer
			}
		}
		proposed[len(current)] = c

		res := s.dfs_expand(proposed, candidates[i+1:])
		if len(res) > len(largest) {
			// largest = slices.Clone(res)
			largest = res
		}
	}
	return largest
}

func (s *solution) run2() {
}

func (s *solution) res() int {
	return len(s.connet3)
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}
	collections := map[string][]string{}
	for _, line := range lines {
		items := strings.Split(line, "-")
		a, b := items[0], items[1]
		collections[a] = append(collections[a], b)
		collections[b] = append(collections[b], a)
	}

	return &solution{
		connet3:   map[string]bool{},
		relations: collections,
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
