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
		res := PartOne(os.Stdin)
		fmt.Println(res)
	case "2":
		res := PartTwo(os.Stdin)
		fmt.Println(res)
	}
}

func PartOne(r io.Reader) int {
	lines, err := readLists(r)
	if err != nil {
		return 0
	}
	var ans int
	for _, line := range lines {
		ans += isSafe(line)
	}
	return ans
}

//	func PartTwo(r io.Reader) int {
//		lines, err := readLists(r)
//		if err != nil {
//			return 0
//		}
//
//		ans, curAns := 0, 0
//
//		for _, line := range lines {
//			curAns = isSafe(line)
//			if curAns == 0 {
//				for i := range line {
//					newLine := []int{}
//					newLine = append(newLine, line[:i]...)
//					newLine = append(newLine, line[i+1:]...)
//					curAns = isSafe(newLine)
//					if curAns == 1 {
//						break
//					}
//				}
//			}
//			ans += curAns
//		}
//
//		return ans
//	}
func PartTwo(r io.Reader) int {
	lines, err := readLists(r)
	if err != nil {
		return 0
	}

	curAns, ans := 0, 0
	for _, line := range lines {
		curAns = isSafe2(line)
		if curAns == 0 {
			curAns = isSafe(line[1:])
		}
		ans += curAns
	}

	return ans
}

func inValid(a, b int) bool {
	return a*b <= 0 || utils.Abs(b) > 3
}

func isSafe(l []int) int {
	initDiff := l[1] - l[0]
	for i := 1; i < len(l); i++ {
		curDiff := l[i] - l[i-1]
		if inValid(initDiff, curDiff) {
			return 0
		}
	}
	return 1
}

func isSafe2(l []int) int {
	skip := 0
	initDiff := l[1] - l[0]
	for i := 1; i < len(l); i++ {
		curDiff := l[i] - l[i-1]
		if skip == 1 {
			curDiff = l[i] - l[i-2]
			skip++
			if i == 2 {
				initDiff = l[2] - l[0]
			}
		}
		if inValid(initDiff, curDiff) {
			if skip == 0 {
				skip++
				continue
			} else {
				return 0
			}
		}
	}
	return 1
}

func readLists(r io.Reader) ([][]int, error) {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		return nil, fmt.Errorf("could not read input: %w", err)
	}

	result := make([][]int, len(lines))
	for i, line := range lines {
		nums := utils.IntsFromString(line)
		if nums == nil {
			return nil, fmt.Errorf("no number found %v", err)
		}
		result[i] = nums
	}
	return result, nil
}
