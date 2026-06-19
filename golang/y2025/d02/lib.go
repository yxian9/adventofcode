package main

import (
	"strconv"
	"strings"
)

type Range struct {
	start, stop, len int
}

func (r Range) InvalidIDs(chunkLen int) map[int]struct{} {
	repeat := r.len / chunkLen

	invalids := make(map[int]struct{})

	for i := r.start / pow10(r.len-chunkLen); i <= r.stop/(pow10(r.len-chunkLen)+1); i++ {
		// for i := r.start / pow10(r.len/2); i <= r.stop; i++ {
		id := buildID(i, repeat)
		if id < r.start {
			continue
		}
		if id > r.stop {
			break
		}
		invalids[id] = struct{}{}
	}

	return invalids
}

func parseInput(s string) []Range {
	var list []Range
	for r := range strings.SplitSeq(s, ",") {
		limits := strings.Split(r, "-")
		start, _ := strconv.Atoi(limits[0])
		stop, _ := strconv.Atoi(limits[1])
		startDigits, stopDigits := intLen(start), intLen(stop)

		// split the range into two ranges
		// e.g., start->999 and 1000->stop
		// to work on single digits range later
		if startDigits == stopDigits {
			list = append(list, Range{start, stop, startDigits})
		} else {
			list = append(list, Range{start, pow10(stopDigits-1) - 1, startDigits})
			list = append(list, Range{pow10(startDigits), stop, stopDigits})
		}
	}
	return list
}

// faster than
// int(math.Pow10(exp))
func pow10(exp int) int {
	result := 1
	for range exp {
		result *= 10
	}
	return result
}

// faster than
// len(strconv.Itoa(i))
func intLen(i int) int {
	// if i == 0 {
	// 	return 1
	// }
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

// faster than
// strconv.Atoi(strings.Repeat(strconv.Itoa(chunk), repeat))
func buildID(chunk, repeat int) int {
	n := chunk
	for i := 1; i < repeat; i++ {
		n = n*pow10(intLen(chunk)) + chunk
	}
	return n
}
