package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type solution struct {
	input      []string
	ans1, ans2 int
}

func buildSolution(r io.Reader) *solution {
	lines, _ := utils.LinesFromReader(r)
	return &solution{
		input: lines,
		ans1:  0,
		ans2:  0,
	}
}

type cmdRes struct {
	cd   bool
	file bool
	path string
	size int
}

func parse(line string) cmdRes {
	items := strings.Fields(line)
	if items[0] == "$" {
		switch items[1] {
		case "cd":
			return cmdRes{cd: true, path: items[2]}
			// case "ls":
			// 	return cmdRes{ls: true}
		}
	}
	if items[0] != "dir" && items[1] != "ls" {
		size, _ := strconv.Atoi(items[0])
		return cmdRes{file: true, size: size, path: items[1]}
	}
	return cmdRes{}
}

func (s *solution) run1() {
	var pathList utils.List[string]

	sizeMap := make(map[string]int)

	for _, line := range s.input {
		cmd := parse(line)

		switch {

		case cmd.cd:
			if cmd.path == ".." {
				pathList.Pop()
			} else {
				pathList.Push(cmd.path)
			}
			fmt.Printf("%q\n", pathList)

		case cmd.file:
			size := cmd.size
			fmt.Println("file \" ", cmd.path, " \" size: ", cmd.size)

			var b strings.Builder

			for i, p := range pathList.Arr {
				if i > 1 {
					b.WriteString("/")
				}
				b.WriteString(p)
				fmt.Println("\"", b.String(), "\" add ", size)
				sizeMap[b.String()] += size
			}
		}
	}
	for _, v := range sizeMap {
		fmt.Println(v)
		if v <= 100000 {
			s.ans1 += v
		}
	}
	fmt.Printf("%v \n", sizeMap)
	totalSize := 70000000
	require := 30000000
	taken := sizeMap["/"]
	curFree := totalSize - taken
	moreFreeNeed := require - curFree
	sizeList := make([]int, len(sizeMap))
	i := 0
	for _, v := range sizeMap {
		sizeList[i] = v
		i++
	}
	slices.Sort(sizeList)
	for _, v := range sizeList {
		if v >= moreFreeNeed {
			s.ans2 = v
			break
		}
	}
}

func (s *solution) run2() {
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
