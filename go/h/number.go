package h

import (
	"fmt"
	"reflect"
)

// IntsFromString returns a slice of integers in str, where these numbers are
// separated by non-number runes. If a dash preceding a number is the only
// non-number rune between two numbers, it is considered a separator; otherwise,
// it is considered part of the number, which is consequently negative.
//
//	func IntsFromString(str string) []int {
//		words := splitStringIntoIntStrings(str)
//
//		ints := make([]int, len(words))
//
//		for i, w := range words {
//			n, err := strconv.Atoi(w)
//			if err != nil {
//				panic(fmt.Sprintf("could not parse int %q: %v", w, err))
//			}
//
//			ints[i] = n
//		}
//
//		return ints
//	}
func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Isdigit(b any) (int, bool) {
	switch v := b.(type) {
	case rune:
		if v >= '0' && v <= '9' {
			return int(v - '0'), true
		}
		return 0, false
	case byte:
		if v >= '0' && v <= '9' {
			return int(v - '0'), true
		}
		return 0, false
	default:
		fmt.Println("found type:", reflect.TypeOf(v))
		panic("in correct type, expect byte and rune")
	}
}

func IsdigitBool(b any) bool {
	_, ok := Isdigit(b)
	return ok
}
func IntsFromString(str string) (intSlice []int) {
	// wordBuf := make([]rune, 0, len(str))
	wordBuf := 0

	n, sign := 0, 1
	for _, r := range str {

		if r == '-' && wordBuf == 0 {
			// wordBuf = append(wordBuf, r)
			sign = -1
			continue
		}

		if r >= '0' && r <= '9' {
			n *= 10
			n += int(r - '0')
			wordBuf++
			continue
		}

		if wordBuf > 0 {
			intSlice = append(intSlice, n*sign)
			n, wordBuf, sign = 0, 0, 1
		}
		sign = 1
	}

	if wordBuf > 0 {
		intSlice = append(intSlice, n*sign)
	}

	return intSlice
}
