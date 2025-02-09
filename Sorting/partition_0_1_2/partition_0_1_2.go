package main

import (
	"fmt"
)

func swap(x, y int) (int, int) { return y, x }

func Partition012(arr []int, size int) {
	left := 0
	right := size - 1
	i := 0
	for i <= right {
		if arr[i] == 0 {
			arr[i], arr[left] = swap(arr[i], arr[left])
			i++
			left++
		} else if arr[i] == 2 {
			arr[i], arr[right] = swap(arr[i], arr[right])
			right--
		} else {
			i++
		}
	}
}

func main() {
	test_arr := []int{0, 1, 2, 0, 1, 2}
	fmt.Println("Original Array : ", test_arr)
	Partition012(test_arr, len(test_arr))
	fmt.Println("Partitioned Array : ", test_arr)
}
