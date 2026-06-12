package sorting

func QuickSort(arr *[]int, left int, right int) {
	if left < right {
		pi := partition(arr, left, right)
		QuickSort(arr, left, pi-1)
		QuickSort(arr, pi+1, right)
	}
}

func partition(arr *[]int, left int, right int) int {
	pivot := (*arr)[right]

	i := left - 1
	for j := left; j < right; j++ {
		if (*arr)[j] <= pivot {
			i += 1
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
		}
	}
	(*arr)[i+1], (*arr)[right] = (*arr)[right], (*arr)[i+1]
	return i + 1
}
