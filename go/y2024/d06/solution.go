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
		res, err := Part1(os.Stdin)
		if err != nil {
			fmt.Println("p1 error ", err)
		}
		fmt.Println("p1 res 🙆-> ", res)
	case "2":
		res, err := Part2(os.Stdin)
		if err != nil {
			fmt.Println("p2 error ", err)
		}
		fmt.Println("p2 res 🙆-> ", res)
	}
}

var dirs = [4][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func buildGrid(lines []string) (seenGrid [][]int, sx, sy int) {
	seenGrid = make([][]int, len(lines))
	for i, line := range lines {
		seenGrid[i] = make([]int, len(line))
		for j, r := range line {
			if r == '^' {
				sx, sy = i, j
			}
		}
	}
	return
}

func Part1(r io.Reader) (int, error) {
	lines, err := readLists(r)
	angle, answer := 0, 0
	seenGrid, sx, sy := buildGrid(lines)
	nrow, ncol := len(seenGrid), len(seenGrid[0])

	var (
		inside func(int, int) bool
		nextR  func(int, int) rune
		dfs    func(int, int, int)
	)

	inside = func(i, j int) bool {
		return i >= 0 && i < nrow && j >= 0 && j < ncol
	}

	nextR = func(x, y int) rune {
		return rune(lines[x][y])
	}

	dfs = func(cx, cy, angle int) {
		seenGrid[cx][cy] = 1
		curDir := dirs[angle%4]
		nx, ny := cx+curDir[0], cy+curDir[1]
		if !inside(nx, ny) {
			return
		}
		if nextR(nx, ny) != '#' {
			dfs(nx, ny, angle)
		} else {
			angle++
			dfs(cx, cy, angle)
		}
	}

	dfs(sx, sy, angle)
	for _, row := range seenGrid {
		for _, v := range row {
			if v == 1 {
				answer++
			}
		}
	}
	// for _, row := range seenGrid {
	// 	fmt.Println(row)
	// }

	if err != nil {
		fmt.Println(lines)
		return 0, fmt.Errorf("error %w", err)
	}

	return answer, nil
}

func buildGrid2(lines []string) (spotGrid [][]int, sx, sy int) {
	spotGrid = make([][]int, len(lines))
	for i, line := range lines {
		spotGrid[i] = make([]int, len(line))
		for j, r := range line {
			if r == '^' {
				sx, sy = i, j
			}
			if r == '.' {
				spotGrid[i][j] = 1
			}
		}
	}
	return
}

func Part2(r io.Reader) (int, error) {
	dirs := [4][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	lines, err := readLists(r)
	angle, answer := 0, 0
	spotGrid, sx, sy := buildGrid2(lines)
	nrow, ncol := len(lines), len(lines[0])

	var (
		inside    func(int, int) bool
		checkNext func(int, int) rune
		isInfinit func(int, int, int, map[string]bool) int
	)

	inside = func(i, j int) bool {
		return i >= 0 && i < nrow && j >= 0 && j < ncol
	}

	checkNext = func(x, y int) rune {
		return rune(lines[x][y])
	}

	isInfinit = func(cx, cy, angle int, seen map[string]bool) int {
		curDir := dirs[angle%4]
		nx, ny := cx+curDir[0], cy+curDir[1]
		if !inside(nx, ny) {
			return 0
		}
		if checkNext(nx, ny) != '#' && spotGrid[nx][ny] != -1 {
			return isInfinit(nx, ny, angle, seen)
		} else {
			coord := fmt.Sprintf("%d,%d,%d", nx, ny, angle%4)
			if seen[coord] {
				return 1
			}
			seen[coord] = true
			angle++
			return isInfinit(cx, cy, angle, seen)
		}
	}

	// isInfinit = func(cx, cy, angle int, seen map[string]bool) int {
	// 	for {
	// 		curDir := dirs[angle%4]
	// 		nx, ny := cx+curDir[0], cy+curDir[1]
	// 		if !inside(nx, ny) {
	// 			return 0
	// 		}
	// 		if checkNext(nx, ny) != '#' && spotGrid[nx][ny] != -1 {
	// 			cx, cy = nx, ny
	// 		} else {
	// 			coord := fmt.Sprintf("%d,%d,%d", nx, ny, angle%4)
	// 			if seen[coord] {
	// 				return 1
	// 			}
	// 			seen[coord] = true
	// 			angle++
	// 		}
	// 	}
	// }

	for i, row := range lines {
		for j := range row {
			if spotGrid[i][j] == 1 {
				seen := map[string]bool{}
				spotGrid[i][j] = -1
				answer += isInfinit(sx, sy, angle, seen)
				spotGrid[i][j] = 1
			}
		}
	}
	// for _, row := range lines {
	// 	fmt.Println(row)
	// }
	// for _, row := range spotGrid {
	// 	fmt.Println(row)
	// }

	if err != nil {
		fmt.Println(lines)
		return 0, fmt.Errorf("error %w", err)
	}

	return answer, nil
}

func readLists(r io.Reader) ([]string, error) {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		return nil, fmt.Errorf("could not read input: %w", err)
	}

	return lines, err
}
