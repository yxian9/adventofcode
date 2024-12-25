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
	relations              map[string][]string
	ans                    int
	result                 map[string]bool
	children               map[string]map[string]bool
	longest_len            int
	part1Count, part2Count int
	path                   []string
	memo                   map[string]int
	path_result            []string
	longest                string
}

func (s *solution) dfs1(prv string) {
	if len(s.path) == 3 {
		first, last := s.path[0], s.path[2]
		if s.children[first][last] {
			for _, v := range s.path {
				if v[0] == 't' {
					s.path_result = append(s.path_result, fmt.Sprint(s.path))
					break
				}
			}
		}
		return
	}
	for k := range s.children {
		if len(s.path) == 0 || (k > prv && s.children[prv][k]) {
			s.path = append(s.path, k)
			s.dfs1(k)
			s.path = s.path[:len(s.path)-1]
		}
	}
	// for k := range s.children {
	// 	if len(s.path) == 0 {
	// 		s.seen[k] = true
	// 		s.path = append(s.path, k)
	// 		s.dfs1(k)
	// 		s.seen[k] = false
	// 		s.path = s.path[:len(s.path)-1]
	// 		continue
	// 	}
	// 	if k > prv {
	// 		if s.children[prv][k] {
	// 			s.path = append(s.path, k)
	// 			s.dfs1(k)
	// 			s.path = s.path[:len(s.path)-1]
	// 		}
	// 	}
	// }
}

func (s *solution) dfs2(prv string, checked int) {
	if len(s.path) > (s.longest_len) {
		cur_loop := strings.Join(s.path, ",")
		s.longest = cur_loop
		s.longest_len = len(s.path)
	}

	remain := len(s.children) - checked
	if s.longest_len > remain+len(s.path) {
		return
	}

	for k := range s.children {
		if len(s.path) == 0 || (k > prv && s.all_connected(k)) {
			s.path = append(s.path, k)
			s.dfs2(k, checked+1)
			s.path = s.path[:len(s.path)-1]
		}
	}
}

func (s *solution) dfs3(prv string) {
	if len(s.path) > (s.longest_len) {
		cur_loop := strings.Join(s.path, ",")
		s.longest = cur_loop
		s.longest_len = len(s.path)
	}

	// if len(s.path) == 0 {
	// 	s.path = append(s.path, prv)
	// 	for k := range s.children[prv] {
	// 		if k > prv && s.all_connected(k) {
	// 			s.path = append(s.path, k)
	// 			s.dfs3(k)
	// 			s.path = s.path[:len(s.path)-1]
	// 		}
	// 	}
	// 	s.path = s.path[:len(s.path)-1]
	// } else {
	for k := range s.children[prv] {
		if k > prv && s.all_connected(k) {
			s.path = append(s.path, k)
			s.dfs3(k)
			s.path = s.path[:len(s.path)-1]
		}
	}
	// }
}

func (s *solution) all_connected(check string) bool {
	for _, v := range s.path {
		if !s.children[v][check] {
			return false
		}
	}
	return true
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
			for j := i + 1; j < len(a_childs); j++ {
				c := a_childs[j] // both b and c need to be a's child
				if slices.Contains(s.relations[b], c) {
					keyslice := []string{a, b, c}
					sort.Strings(keyslice)
					key := fmt.Sprint(keyslice)
					s.result[key] = true
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
	fmt.Println("part1 longest", strings.Join(largest, ","))
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
	s.dfs1("")
	// fmt.Println("part2", s.path_result)
	for k := range s.children {
		s.path = append(s.path, k)
		s.dfs3(k)
		s.path = s.path[:len(s.path)-1]
	}
	fmt.Println("part2 longest", s.longest)
	// for a, a_childs := range s.relations {
	// 	if len(a_childs) < 2 {
	// 		continue
	// 	}
	// 	for _, b := range a_childs {
	// 		if a > b {
	// 			continue
	// 		}
	// 		for _, c := range a_childs {
	// 			// c := a_childs[j] // both b and c need to be a's child
	// 			if b > c {
	// 				continue
	// 			}
	// 			if slices.Contains(s.relations[b], c) {
	// 				if a[0] == 't' || b[0] == 't' || c[0] == 't' {
	//
	// 					keyslice := []string{a, b, c}
	// 					fmt.Println("part1, i, j, c:", a, b, c)
	// 					sort.Strings(keyslice)
	// 					key := fmt.Sprint(keyslice)
	// 					s.result[key] = true
	// 				}
	// 			}
	// 		}
	// 	}
	// }
	//
	// var largest []string
	// for a, v := range s.relations {
	// 	cohort := s.dfs_expand([]string{a}, v)
	// 	if len(cohort) > len(largest) {
	// 		largest = cohort
	// 	}
	// }
	// slices.Sort(largest)
	// // fmt.Println(strings.Join(largest, ","))
}

func (s *solution) res() int {
	return len(s.result)
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}
	collections := map[string][]string{}
	children := map[string]map[string]bool{}
	for _, line := range lines {
		items := strings.Split(line, "-")

		// a, b := items[0], items[1]
		// collections[a] = append(collections[a], b)
		// collections[b] = append(collections[b], a)
		for i := range 2 {
			a, b := items[i], items[1-i]
			collections[a] = append(collections[a], b)
			items := children[a]
			if items == nil {
				items = map[string]bool{}
				children[a] = items
			}
			items[b] = true
		}

	}

	return &solution{
		result:    map[string]bool{},
		relations: collections,
		children:  children,
		memo:      map[string]int{},
		path:      make([]string, 0, len(children)),
	}
}

func part1(r io.Reader) int {
	s := buildSolution(r)
	s.run1()
	fmt.Println("part1 count", s.part1Count)
	return s.res()
}

func part2(r io.Reader) int {
	s := buildSolution(r)
	s.run2()
	fmt.Println("part2 count", s.part2Count)
	return len(s.path_result)
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
