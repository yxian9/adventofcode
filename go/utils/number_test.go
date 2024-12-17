package utils

import (
	"reflect"
	"testing"
)

func TestSplitStringIntoIntStrings(t *testing.T) {
	testCases := []struct {
		str  string
		want []int
	}{
		{
			str:  "1 23 4 567 8 90",
			want: []int{1, 23, 4, 567, 8, 90},
		},
		{
			str:  "1,23,4,567,8,90",
			want: []int{1, 23, 4, 567, 8, 90},
		},
		{
			str:  "1 23,4,   567+8//90",
			want: []int{1, 23, 4, 567, 8, 90},
		},
		{
			str:  "1234567890",
			want: []int{1234567890},
		},
		{
			str:  "hello world",
			want: nil,
		},
		{
			str:  "---",
			want: nil,
		},
		{
			str:  "- 1 -1 2",
			want: []int{1, -1, 2},
		},
		{
			str:  "-000",
			want: []int{0},
		},
		{
			str:  "-1 23 -4 567 -890",
			want: []int{-1, 23, -4, 567, -890},
		},
		{
			str:  "-1 23-4 567-890",
			want: []int{-1, 23, 4, 567, 890},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.str, func(t *testing.T) {
			got := IntsFromString(tc.str)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}
}
