package main

import (
	"adventofcode/golang/h"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"time"
)

type solution struct {
	input      []string
	ans1, ans2 int
}

func buildSolution(r io.Reader) *solution {
	lines, _ := h.LinesFromReader(r)
	return &solution{
		input: lines,
		ans1:  0,
		ans2:  0,
	}
}

type T struct {
	dst, src, len int
}

func parse(l []string) (seed []int, relation [][]T) {
	for i := 0; i < len(l); i++ {
		if i == 0 {
			_, remain, _ := strings.Cut(l[i], ":")
			seed = h.IntsFromString(remain)
			i++
			continue
		}

		var cur = []T{}
		for ; i < len(l) && l[i] != ""; i++ {
			ints := h.IntsFromString(l[i])
			if len(ints) != 0 {
				cur = append(cur, T{ints[0], ints[1], ints[2]})
			}
		}
		relation = append(relation, cur)
	}
	return seed, relation
}

func transform(seed []int, relation []T) (res []int) {
top:
	for _, i := range seed {
		for _, t := range relation {
			if i >= t.src && i < t.src+t.len {
				res = append(res, t.dst+i-t.src)
				continue top
			}
		}
		res = append(res, i)
	}
	return
}
func transform2(seed int, relation []T) (res int) {
	for _, t := range relation {
		if seed >= t.src && seed < t.src+t.len {
			return t.dst + seed - t.src
		}
	}
	return seed
}

func (s *solution) run1() {
	seed, relations := parse(s.input)
	// fmt.Println(seed)
	// for _, v := range relations {
	// 	fmt.Println(v)
	// }
	for _, relation := range relations {
		seed = transform(seed, relation)
		// fmt.Println(seed)
	}
	s.ans1 = slices.Min(seed)
}

func worker(s, step int, relations [][]T, minChan chan<- int) {
	res := 1<<32 - 1
	v := 0
	for j := range step {
		v = s + j
		for _, relation := range relations {
			v = transform2(v, relation)
		}
		res = min(res, v)
	}
	minChan <- res
}

func (s *solution) run2() {
	seed, relations := parse(s.input)
	s.ans2 = 1<<32 - 1
	var minChan = make(chan int)
	for i := 0; i < len(seed); i = i + 2 {
		go worker(seed[i], seed[i+1], relations, minChan)
	}
	for i := 0; i < len(seed); i = i + 2 {
		s.ans2 = min(s.ans2, <-minChan)
	}

}

func (s *solution) res1() int {
	return s.ans1
}

func (s *solution) res2() int {
	return s.ans2
}

func part1(r io.Reader) int {
	s := buildSolution(r)
	s.run1()
	return s.res1()
}

func part2(r io.Reader) int {
	s := buildSolution(r)
	s.run2()
	return s.res2()
}

func main() {

	_, filename, _, _ := runtime.Caller(0)
	Input, err := os.Open(filepath.Join(filepath.Dir(filename), "input.txt"))
	if err != nil {
		log.Fatalf("fail open input.txt %v", err)
	}
	start := time.Now()
	result := part1(Input)
	elapsed := time.Since(start)
	fmt.Printf("p1 res 🙆-> %d (Time taken: %s)\n", result, elapsed)
	Input.Close()
	Input, err = os.Open(filepath.Join(filepath.Dir(filename), "input.txt"))
	start = time.Now()
	result = part2(Input)
	elapsed = time.Since(start)
	fmt.Printf("p2 res 🙆-> %d (Time taken: %s)\n", result, elapsed)
	Input.Close()
}
