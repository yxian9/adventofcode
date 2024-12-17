package main

import (
	"strings"
	"testing"
)

const input1 = `
7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`

func TestPart1(t *testing.T) {
	r := strings.NewReader(input1)
	got := PartOne(r)
	want := 2

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	r := strings.NewReader(input1)
	got := PartTwo(r)
	want := 4

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
