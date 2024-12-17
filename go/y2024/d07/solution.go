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

func Part1(r io.Reader) (int, error) {
	lines, err := readLists(r)
	if err != nil {
		fmt.Println(lines)
		return 0, fmt.Errorf("error %w", err)
	}
	ans := 0
	for _, line := range lines {
		intslice := utils.IntsFromString(line)
		if isvalid(intslice, ops) {
			ans += intslice[0]
		}
	}
	return ans, nil
}

var ops = []func(int, int) int{
	func(a, b int) int {
		return a + b
	},
	func(a, b int) int {
		return a * b
	},
}

var ops2 = append(ops,
	func(a, b int) int {
		factor := 1
		for factor <= b {
			factor *= 10
		}
		return a*factor + b
	})

func isvalid(intslice []int, ops []func(int, int) int) bool {
	target, input := intslice[0], intslice[1:]

	var dfs func(int, int, []int) bool
	dfs = func(level, curTotal int, input []int) bool {
		// if curTotal == target && level == len(input) {
		// 	return true
		// }
		// if curTotal > target || level > len(input)-1 {
		// 	return false
		// }

		if level == len(input) {
			return curTotal == target
		}
		if curTotal > target {
			return false
		}

		for _, operator := range ops {
			if dfs(level+1, operator(curTotal, input[level]), input) {
				return true
			}
		}

		return false
	}

	return dfs(1, input[0], input)
}

func Part2(r io.Reader) (int, error) {
	lines, err := readLists(r)
	if err != nil {
		fmt.Println(lines)
		return 0, fmt.Errorf("error %w", err)
	}
	ans := 0
	for _, line := range lines {
		intslice := utils.IntsFromString(line)
		if isvalid(intslice, ops2) {
			ans += intslice[0]
		}
	}
	return ans, nil
}

func readLists(r io.Reader) ([]string, error) {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		return nil, fmt.Errorf("could not read input: %w", err)
	}

	return lines, err
}
