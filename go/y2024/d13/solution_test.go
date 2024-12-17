package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	testInput, err := os.Open("test1.txt")
	if err != nil {
		log.Fatalf("fail open test1.txt %v", err)
	}
	want := 480

	t.Run(fmt.Sprintf("want: %d", want), func(t *testing.T) {
		got := part2(testInput)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestPart2(t *testing.T) {
	testInput, err := os.Open("test1.txt")
	if err != nil {
		log.Fatalf("fail open test1.txt %v", err)
	}
	want := 0

	t.Run(fmt.Sprintf("want: %d", want), func(t *testing.T) {
		got := part2(testInput)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
