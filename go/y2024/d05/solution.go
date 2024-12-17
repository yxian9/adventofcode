package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"os"
	"slices"
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

func readLists(r io.Reader) ([]string, error) {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		return nil, fmt.Errorf("could not read input: %w", err)
	}

	return lines, err
}

type (
	manual   []int
	ajacList map[int][]int
)

func buildAjacAndManuals(r io.Reader) (ajacList, []manual, error) {
	lines, err := readLists(r)
	if err != nil {
		return nil, nil, fmt.Errorf("read error %w", err)
	}
	depences, manuals := make(ajacList), make([]manual, 0)
	processAjac := true
	for _, line := range lines {
		nums := utils.IntsFromString(line)
		if len(nums) == 0 {
			processAjac = false
			continue
		}
		if processAjac {
			if len(nums) != 2 {
				return nil, nil, fmt.Errorf("length incorrect")
			}
			depences[nums[0]] = append(depences[nums[0]], nums[1])
		} else {
			manuals = append(manuals, nums)
		}
	}
	return depences, manuals, nil
}

func (m manual) median() int {
	return m[len(m)/2]
}

func (m manual) isSorted(ajac ajacList) bool {
	seen := map[int]bool{}

	for _, v := range m {
		for _, after := range ajac[v] {
			if seen[after] {
				return false
			}
		}
		seen[v] = true
	}
	return true
}

func (m manual) getLocalAjc(ajac ajacList) ajacList {
	localAjac := make(ajacList)
	for _, v := range m {
		filteredDep := []int{}
		for _, num := range ajac[v] {
			if slices.Contains(m, num) {
				filteredDep = append(filteredDep, num)
			}
		}
		localAjac[v] = filteredDep
	}
	return localAjac
}

func (m manual) Sort(ajac ajacList) int {
	localAjac := m.getLocalAjc(ajac)

	indegree := map[int]int{}
	for _, depences := range localAjac {
		for _, num := range depences {
			indegree[num]++
		}
	}

	topSort, queue := []int{}, []int{}

	for _, v := range m {
		if indegree[v] == 0 {
			queue = append(queue, v)
		}
	}
	for len(queue) > 0 {
		header := queue[0]
		queue = queue[1:]
		topSort = append(topSort, header)
		for _, num := range localAjac[header] {
			indegree[num]--
			if indegree[num] == 0 {
				queue = append(queue, num)
			}
		}
	}
	return topSort[len(topSort)/2]
}

func Part1(r io.Reader) (int, error) {
	ajacList, manuals, err := buildAjacAndManuals(r)
	if err != nil {
		return 0, fmt.Errorf("error %w", err)
	}
	ans := 0
	for _, manual := range manuals {
		if manual.isSorted(ajacList) {
			ans += manual.median()
		}
	}

	return ans, nil
}

func Part2(r io.Reader) (int, error) {
	ajacList, manuals, err := buildAjacAndManuals(r)
	if err != nil {
		return 0, fmt.Errorf("error %w", err)
	}
	ans := 0
	for _, manual := range manuals {
		if !manual.isSorted(ajacList) {
			ans += manual.Sort(ajacList)
		}
	}

	return ans, nil
}
