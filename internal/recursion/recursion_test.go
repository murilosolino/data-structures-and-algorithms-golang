package recursion_test

import (
	"testing"

	"github.com/murilosolino/data-structures-and-algorithms/internal/recursion"
)

func TestSum(t *testing.T) {
	arr := []int{2, 4, 6}
	r := recursion.Sum(arr)
	if r != 12 {
		t.Fatalf("expected %v, recived: %v", 12, r)
	}
}

func TestCountItens(t *testing.T) {
	arr := []int{2, 4, 6}
	r := recursion.CountItens(arr)
	if r != len(arr) {
		t.Fatalf("expected %v, recived: %v", 3, r)
	}
}

func TestMax(t *testing.T) {
	arr := []int{2, 4, 115, 6, 5, 23, 77, 34, 10}
	r := recursion.Max(arr)
	if r != 115 {
		t.Fatalf("expected %v, recived: %v", 115, r)
	}
}
