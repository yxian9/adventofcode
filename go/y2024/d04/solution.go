package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"os"
)

func main() {
	arg := os.Args[1]
	fmt.Println("Running part", arg)
	switch arg {
	case "1":
		res := Part1(os.Stdin)
		fmt.Println(res)
	case "2":
		res := Part2(os.Stdin)
		fmt.Println(res)
	}
}

const XMAS = "XMAS"

func Part1(r io.Reader) int {
	lines, err := readLists(r)
	if err != nil {
		return 0
	}
	var dirs [][2]int
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			dirs = append(dirs, [2]int{i, j})
		}
	}
	ans := 0
	nrow, ncol := len(lines), len(lines[3])
	var inside func(i, j int) bool
	inside = func(i, j int) bool {
		return i >= 0 && i < nrow && j >= 0 && j < ncol
	}

	for i, line := range lines {
		for j := range line {
			// if c != 'X' {
			// 	continue
			// }
		toploop:
			for _, dir := range dirs {
				// valid := true

				for step, CheckRune := range XMAS {
					nx, ny := i+dir[0]*step, j+dir[1]*step
					if !inside(nx, ny) || CheckRune != rune(lines[nx][ny]) {
						// valid = false
						// break
						continue toploop
					}
				}
				// if valid {
				ans++
				// }
			}
		}
	}
	return ans
}

func Part2(r io.Reader) int {
	lines, err := readLists(r)
	if err != nil {
		return 0
	}
	dirs := [2][2][2]int{
		{
			{1, 1},
			{-1, -1},
		},
		{
			{-1, 1},
			{1, -1},
		},
	}

	valueM := map[rune]int{
		'M': 1,
		'S': 2,
	}

	ans := 0
	nrow, ncol := len(lines), len(lines[0])
	var inside func(i, j int) bool
	inside = func(i, j int) bool {
		return i >= 0 && i < nrow && j >= 0 && j < ncol
	}
	for i, line := range lines {
	toploop:
		for j, c := range line {
			if c != 'A' {
				continue
			}
			// valid := true

			for _, dir := range dirs {
				curCount := 0
				for _, point := range dir {
					nx, ny := i+point[0], j+point[1]
					if !inside(nx, ny) {
						// valid = false
						continue toploop
					}
					curCount += valueM[rune(lines[nx][ny])]
				}
				if curCount != 3 {
					// valid = false
					continue toploop
				}
			}
			// if valid {
			ans++
			// }
		}
	}
	return ans
}

func readLists(r io.Reader) ([]string, error) {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		return nil, fmt.Errorf("could not read input: %w", err)
	}

	return lines, nil
}
