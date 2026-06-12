package sorting_test

import (
	"reflect"
	"testing"

	"github.com/murilosolino/data-structures-and-algorithms/internal/sorting"
)

func TestBubbleSort(t *testing.T) {
	var arr = []int{5, 4, 2, 3, 1}
	var sortArr = []int{1, 2, 3, 4, 5}
	sorting.BubbleSort(&arr)
	if !reflect.DeepEqual(arr, sortArr) {
		t.Errorf("falha na ordenação.\nesperado: %v \nrecebido: %v", sortArr, arr)
	}
}

func TestQuickSort(t *testing.T) {
	var arr = []int{5, 4, 2, 3, 1}
	var sortArr = []int{1, 2, 3, 4, 5}
	sorting.QuickSort(&arr, 0, len(arr)-1)
	if !reflect.DeepEqual(arr, sortArr) {
		t.Errorf("falha na ordenação.\nesperado: %v \nrecebido: %v", sortArr, arr)
	}
}
