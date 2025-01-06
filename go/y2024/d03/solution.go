package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type solution struct {
	input string
	ans   int
}

const (
	mul     = "mul("
	doStr   = "do()"
	dontStr = "don't()"
)

func getNum(input string, s int) (res, length int) {
	for i := s; i < len(input); i++ {
		cha := rune(input[i])
		if cha >= '0' && cha <= '9' {
			res *= 10
			res += int(cha - '0')
			length++
		} else {
			return res, length
		}
	}
	return res, length
}

func (s *solution) run1(input string) {
	for idx := 0; idx < len(input)-4; idx++ {
		if input[idx:idx+len(mul)] != mul {
			// idx += len(mul) - 1
			continue
		}
		// start count number
		idx += len(mul)
		p1, p1Len := getNum(input, idx)
		idx += p1Len
		if input[idx] != ',' {
			continue
		}
		idx++ // move over ,
		p2, p2Len := getNum(input, idx)
		idx += p2Len
		if input[idx] != ')' {
			continue
		}
		s.ans += p1 * p2
	}
}

func (s *solution) run2(input string) {
	do := true
	for idx := 0; idx < len(input)-4; idx++ {
		if input[idx:idx+len(doStr)] == doStr {
			do = true
			// idx += len(doStr) - 1 // continue will add one more
			continue
		}
		if idx+len(dontStr) < len(input) && input[idx:idx+len(dontStr)] == dontStr {
			do = false
			// idx += len(dontStr) - 1
			continue
		}
		if input[idx:idx+len(mul)] != mul {
			// idx += len(mul) - 1 // negative selection does not work
			continue
		}
		// start count number
		idx += len(mul)
		p1, p1Len := getNum(input, idx)
		idx += p1Len
		if input[idx] != ',' {
			continue
		}
		idx++ // move over ,
		p2, p2Len := getNum(input, idx)
		idx += p2Len
		if input[idx] != ')' {
			continue
		}
		if do {
			s.ans += p1 * p2
		}
	}
}

func (s *solution) res() int {
	return s.ans
}

func buildSolution(r io.Reader) *solution {
	line, err := utils.ByteSFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", line, err)
	}

	return &solution{
		input: string(line),
		ans:   0,
	}
}

func part1(r io.Reader) int {
	s := buildSolution(r)
	s.run1(s.input)
	return s.res()
}

func part2(r io.Reader) int {
	s := buildSolution(r)
	s.run2(s.input)
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
	Input.Close()
	Input, _ = os.Open("input.txt")
	start = time.Now()
	result = part2(Input)
	elapsed = time.Since(start)
	fmt.Printf("p2 res 🙆-> %d (Time taken: %s)\n", result, elapsed)
}
