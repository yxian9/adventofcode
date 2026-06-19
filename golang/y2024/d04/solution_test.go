package main

import (
	"strings"
	"testing"
)

const input1 = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`

func TestPart1(t *testing.T) {
	r := strings.NewReader(input1)
	got := Part1(r)
	want := 18

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	r := strings.NewReader(input1)
	got := Part2(r)
	want := 9

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
