package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

var t1, t2 = 1227775554, 4174379265

func TestPart1(t *testing.T) {
	testInput, err := os.Open("test1.txt")
	if err != nil {
		log.Fatalf("fail open test1.txt %v", err)
	}

	t.Run(fmt.Sprintf("want: %d", t1), func(t *testing.T) {
		got := part1(testInput)

		if got != t1 {
			t.Errorf("got %v want %v", got, t1)
		}
	})
}

func TestPart2(t *testing.T) {
	testInput, err := os.Open("test1.txt")
	if err != nil {
		log.Fatalf("fail open test1.txt %v", err)
	}

	t.Run(fmt.Sprintf("want: %d", t2), func(t *testing.T) {
		got := part2(testInput)

		if got != t2 {
			t.Errorf("got %v want %v", got, t2)
		}
	})
}
