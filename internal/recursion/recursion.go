package recursion

import "fmt"

func Sum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return arr[0] + Sum(arr[1:])
}

func CountItens(list []int) int {
	if len(list) == 0 {
		return 0
	}
	return 1 + CountItens(list[1:])
}

func Max(list []int) int {
	if len(list) == 0 {
		return 0
	}
	max := list[0]
	v := Max(list[1:])
	if max >= v {
		return max
	}
	return v
}

func Regressiva(i int) {
	if i < 0 {
		return
	}
	fmt.Println(i)
	Regressiva(i - 1)
}
