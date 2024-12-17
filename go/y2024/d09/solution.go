package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type solution struct {
	input    []byte
	intSlice []int
	ans      int
}

func newSlice(fill, length int) (s []int) {
	s = make([]int, length)
	for i := range s {
		s[i] = fill
	}
	return
}

func (s *solution) intSliceFromInput() {
	fId := 0
	for i, b := range s.input {
		if b == '\n' {
			break
		}
		curLength := int(b - '0')
		if i%2 == 0 {
			s.intSlice = append(s.intSlice, newSlice(fId, curLength)...)
			fId++
		} else {
			s.intSlice = append(s.intSlice, newSlice(-1, curLength)...)
		}
	}
	// fmt.Println(s.intSlice)
}

func (s *solution) reorderSlice() {
	// 011111..111....22222..
	//         ^          ^
	// 022111222......
	l, r := 0, len(s.intSlice)-1
	for l < r {
		if s.intSlice[l] != -1 {
			l++
			continue
		}
		if s.intSlice[r] == -1 {
			r--
			continue
		}
		// swap
		s.intSlice[l], s.intSlice[r] = s.intSlice[r], s.intSlice[l]
		l++
		r--
		// continue
		// if s.intSlice[l] == -1 && s.intSlice[r] != -1 {
		// 	l++
		// 	r++
		// 	continue
		// }
	}
	// fmt.Println(s.intSlice)
}

func (s *solution) FindFirstEmpSpot(requiredLen, rightBound int) (int, bool) {
	curLen, start := 0, -1
	for i := 0; i < rightBound; i++ {

		if s.intSlice[i] == -1 {
			curLen++
			if start == -1 {
				start = i
			}
			if curLen == requiredLen {
				return start, true
			}
			continue
		}
		curLen, start = 0, -1
	}
	return -1, false
}

func (s *solution) getFBlockIdxs(fid int) (fbloc []int) {
	for i, v := range s.intSlice {
		if v == fid {
			fbloc = append(fbloc, i)
		}
		if v != fid && len(fbloc) > 0 {
			break
		}
	}
	return fbloc
}

func (s *solution) getFBlockIdx(fid int) (start, len int) {
	for i, v := range s.intSlice {
		if v == fid {
			start = i
			break
		}
	}

	for i, v := range s.intSlice {
		if i < start {
			continue
		}
		if v != fid {
			break
		}
		len++
	}
	return start, len
}

func (s *solution) reorderSlice2() {
	// from back to first build the fileBlocks
	for fID := s.intSlice[len(s.intSlice)-1]; fID >= 0; fID-- {

		// fBlockIdxs := s.getFBlockIdxs(fID)
		fbStart, blockLen := s.getFBlockIdx(fID)

		// lstart, find := s.FindFirstEmpSpot(len(fBlockIdxs), fBlockIdxs[0])
		lstart, find := s.FindFirstEmpSpot(blockLen, fbStart)

		if find {
			// swap
			// for i := 0; i < blockLen; i++ {
			for i := range blockLen {
				s.intSlice[fbStart+i] = -1
				s.intSlice[lstart+i] = fID
			}
		}
	}
}

func (s *solution) reorderSlice3() {
	// 011111..111....22222..
	//       ^            ^
	// 022111222......

	// from back to first build the fileBlocks
	for fID := s.intSlice[len(s.intSlice)-1]; fID >= 0; fID-- {
		// fBlockLen := 0
		fBlockIdxs := []int{}
		for i := len(s.intSlice) - 1; i >= 0; i-- {
			if s.intSlice[i] == fID {
				fBlockIdxs = append(fBlockIdxs, i)
			}
			// else {
			// 	break // break does not work
			// }
		}
		// find emtspot
		empSpotLen, lstart := 0, -1
		// for i, v := range s.intSlice {
		for i := 0; i < fBlockIdxs[0]; i++ {
			if s.intSlice[i] != -1 {
				empSpotLen, lstart = 0, -1
				continue
			}
			if lstart == -1 {
				lstart = i
			}
			empSpotLen++
			if empSpotLen == len(fBlockIdxs) {
				break
			}
		}

		if empSpotLen == len(fBlockIdxs) {
			// swap
			for i, v := range fBlockIdxs {
				s.intSlice[v] = -1
				s.intSlice[lstart+i] = fID
			}
		}
	}
}

func (s *solution) run1() {
	s.intSliceFromInput()
	s.reorderSlice()
}

func (s *solution) run2() {
	s.intSliceFromInput()
	fmt.Println(s.intSlice)
	s.reorderSlice2()
}

func (s *solution) res() int {
	for i, v := range s.intSlice {
		if v == -1 {
			continue
		}
		s.ans += i * v
	}
	return s.ans
}

func buildSolution(r io.Reader) *solution {
	line, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	return &solution{
		input: line,
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
	return s.res()
}

func main() {
	arg := os.Args[1]
	fmt.Println("Running part", arg)
	switch arg {
	case "1":
		fmt.Println("p1 res 🙆-> ", part1(os.Stdin))
	case "2":
		fmt.Println("p2 res 🙆-> ", part2(os.Stdin))
	}
}

//

func byteconvert() {
	a := int('0')
	b := '\n'
	b1 := byte('\n')
	c := int(b - '0')
	e := int(b1 - '0')
	// f := int(b - a)

	fmt.Println(a, b, b1, c, e)
}
