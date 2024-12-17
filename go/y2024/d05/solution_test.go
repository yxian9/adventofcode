package main

import (
	"log"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	testInput, err := os.Open("test1.txt")
	if err != nil {
		log.Fatalf("fail open test1.txt %v", err)
	}

	got, _ := Part1(testInput)
	want := 143

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	testInput, err := os.Open("test1.txt")
	if err != nil {
		log.Fatalf("fail open test1.txt %v", err)
	}
	got, _ := Part2(testInput)
	want := 123

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
