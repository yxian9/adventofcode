package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func (s *solution) run1() {
	s.buildArray(s.input)
	// two point
	l, r := 0, len(s.arr)-1
	for l < r {
		if s.arr[l] != -1 {
			l++
			continue
		}
		if s.arr[r] == -1 {
			r--
			continue
		}
		s.swap(l, r)
		l++
		r--
	}
}

func (s *solution) swap(l, r int) {
	s.arr[l], s.arr[r] = s.arr[r], s.arr[l]
}

func (s *solution) buildArray(input string) {
	var arr []int
	id := 0
	for i, r := range input {
		if i%2 == 0 {
			for range int(r - '0') {
				arr = append(arr, id)
			}
			id++
		} else {
			for range int(r - '0') {
				arr = append(arr, -1)
			}
		}
	}
	s.arr = arr
}

func (s *solution) run2() {
	s.buildArray(s.input)
	fileID := s.arr[len(s.arr)-1]
	for ; fileID >= 0; fileID-- {
		// scan for free space
		chunkL, chunkLen := s.searchFile(fileID)
		freeL, found := s.search(chunkLen, chunkL)
		if found {
			s.swap2(freeL, chunkL, chunkLen, fileID)
		}
	}
}

func (s *solution) searchFile(fildID int) (left int, length int) {
	for i, v := range s.arr {
		if v == fildID {
			if length == 0 {
				left = i
			}
			length++
		}
	}
	return left, length
}

func (s *solution) swap2(freeL, chunkL, chunkLen, fileID int) {
	for i := freeL; i < freeL+chunkLen; i++ {
		s.arr[i] = fileID
	}
	for i := chunkL; i < chunkL+chunkLen; i++ {
		s.arr[i] = -1
	}
}

func (s *solution) search(length, rightB int) (freeL int, found bool) {
	freeSpace := 0
	for i, v := range s.arr {
		if i == rightB {
			break
		}
		if v == -1 {
			if freeSpace == 0 {
				freeL = i
			}
			freeSpace++
			if freeSpace == length {
				return freeL, true
			}
			continue
		}
		freeSpace = 0
	}

	return 0, false
}

func (s *solution) res() int {
	for i, v := range s.arr {
		if v == -1 {
			break
		}
		s.ans += i * v
	}
	return s.ans
}

func (s *solution) res2() int {
	for i, v := range s.arr {
		if v == -1 {
			continue
		}
		s.ans += i * v
	}
	return s.ans
}

type solution struct {
	input string
	arr   []int
	ans   int
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}

	return &solution{
		input: lines[0],
		ans:   0,
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
