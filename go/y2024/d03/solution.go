package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
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

func Part1(r io.Reader) int {
	lines, err := readLists(r)
	if err != nil {
		log.Fatal(err)
	}
	answer := 0
	for _, line := range lines {
		answer += mul(line)
	}
	return answer
}

// func Part2(r io.Reader) int {
// 	lines, err := readLists(r)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	answer := 0
// 	for _, line := range lines {
// 		answer += mul2(line)
// 	}
// 	return answer
// }

func Part2(r io.Reader) int {
	lines, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	return mul3(lines)
}

func getnumber(idx *int, line string) int {
	num := 0
	for *idx < len(line) && line[*idx] >= '0' && line[*idx] <= '9' {
		num = num*10 + int(line[*idx]-'0')
		*idx++
	}
	return num
}

func mul(line string) int {
	// mul(11,8)
	answer := 0
	for i := 0; i < len(line)-4; i++ {
		if line[i:i+4] != "mul(" {
			continue
		}
		i += 4
		x1 := getnumber(&i, line)
		if line[i] != ',' {
			continue
		}
		i++
		x2 := getnumber(&i, line)
		if line[i] != ')' {
			continue
		}
		answer += x1 * x2

	}
	return answer
}

func mul2(line string) int {
	// mul(11,8)
	// don't() do()
	answer := 0
	n := len(line)
	enable := true
	for i := 0; i < n-7; i++ {
		if line[i:i+4] == "do()" {
			enable = true
		}
		if line[i:i+7] == "don't()" {
			enable = false
		}
		if line[i:i+4] == "mul(" {
			i += 4
			x1 := getnumber(&i, line)
			if line[i] != ',' {
				continue
			}
			i++
			x2 := getnumber(&i, line)
			if line[i] != ')' {
				continue
			}
			if enable {
				answer += x1 * x2
			}
		}

	}
	return answer
}

func mul3(input []byte) int {
	// mul(11,8)
	// don't() do()
	answer := 0
	line := string(input)
	n := len(line)
	enable := true
	for i := 0; i < n-7; i++ {
		if line[i:i+4] == "do()" {
			enable = true
		}
		if line[i:i+7] == "don't()" {
			enable = false
		}
		if line[i:i+4] == "mul(" {
			i += 4
			x1 := getnumber(&i, line)
			if line[i] != ',' {
				continue
			}
			i++
			x2 := getnumber(&i, line)
			if line[i] != ')' {
				continue
			}
			if enable {
				answer += x1 * x2
			}
		}

	}
	return answer
}

func readLists(r io.Reader) ([]string, error) {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		return nil, fmt.Errorf("could not read input: %w", err)
	}

	return lines, err
}

func PartTwo(r io.Reader) int {
	corruptedMemory, err := io.ReadAll(r)
	if err != nil {
		return 0
	}

	sum, err := scanCorruptedMemoryWithConditionals(corruptedMemory)
	if err != nil {
		return 0
	}
	return sum
}

var (
	mulInstructionPattern = regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	anyInstructionPattern = regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)|do\(\)|don't\(\)`)
)

func computeMuls(memory []byte) (int, error) {
	matches := mulInstructionPattern.FindAllSubmatch(memory, -1)

	sum := 0
	for _, match := range matches {
		if len(match) != 3 {
			return 0, fmt.Errorf("invalid match: %v", match)
		}

		a := bytesToInt(match[1])
		b := bytesToInt(match[2])

		sum += a * b
	}

	return sum, nil
}

func scanCorruptedMemoryWithConditionals(memory []byte) (int, error) {
	matches := anyInstructionPattern.FindAllSubmatch(memory, -1)

	sum := 0
	doing := true

	for _, match := range matches {
		switch string(match[0]) {
		case "do()":
			doing = true
		case "don't()":
			doing = false
		default:
			if !doing {
				continue
			}
		}

		if len(match) != 3 {
			return 0, fmt.Errorf("invalid match: %v", match)
		}

		a := bytesToInt(match[1])
		b := bytesToInt(match[2])

		sum += a * b
	}

	return sum, nil
}

func bytesToInt(digits []byte) int {
	n := 0
	for _, d := range digits {
		n *= 10
		n += int(d - '0')
	}
	return n
}
