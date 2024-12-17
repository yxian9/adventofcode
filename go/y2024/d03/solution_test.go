package main

import (
	"strings"
	"testing"
)

const input1 = `
xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))
`

const input2 = `
xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))
	`

func TestPart1(t *testing.T) {
	r := strings.NewReader(input1)
	got := Part1(r)
	want := 161

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	r := strings.NewReader(input2)
	got := Part2(r)
	want := 48

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
	// 70478672
}
