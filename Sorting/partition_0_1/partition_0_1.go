package main

import (
	"fmt"
)

func swap(x, y int) (int, int) { return y, x }
func Partition01(arr []int, size int) int {
	left := 0
	right := size - 1
	count := 0
	for left < right {
		for arr[left] == 0 {
			left++
		}
		for arr[right] == 1 {
			right--
		}
		if left < right {
			arr[left], arr[right] = swap(arr[left], arr[right])
			count += 1
		}
	}
	return count
}
func main() {
	test_arr := []int{0, 1, 0, 1, 1, 1}
	fmt.Println("Array before partitioning: ", test_arr)
	size := len(test_arr)
	fmt.Println("Number of swaps required to partition the array: ", Partition01(test_arr, size))
	fmt.Println("Array after partitioning: ", test_arr)
}
