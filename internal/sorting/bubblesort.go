package sorting

func BubbleSort(arr *[]int) {
	size := len(*arr)
	for i := 0; i < size-1; i++ {
		swapped := false
		for j := 0; j < size-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}
