package binarysearch_test

import (
	"testing"

	"github.com/murilosolino/data-structures-and-algorithms/internal/binarysearch"
)

var list = []int{10, 30, 50, 70, 90, 110, 130}
var flagtests = []struct {
	name        string
	arr         []int
	target      int
	targetIndex int
}{
	{
		name:        "1",
		arr:         list,
		target:      30,
		targetIndex: 1,
	},
	{
		name:        "2",
		arr:         list,
		target:      110,
		targetIndex: 5,
	},
	{
		name:        "3",
		arr:         list,
		target:      200,
		targetIndex: -1,
	},
}

func TestBinarySearch(t *testing.T) {
	for _, tt := range flagtests {
		t.Run(tt.name, func(t *testing.T) {
			i := binarysearch.BinarySearch(tt.arr, tt.target)
			if i != tt.targetIndex {
				t.Errorf("expected: %v bute recive: %v", tt.targetIndex, i)
			}
		})
	}
}
