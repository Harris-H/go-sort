package sort

import "go-sort/constraints"

func Bubble[T constraints.Ordered](arr []T) []T {
	n := len(arr)
	var sorted bool // flag to check if the array is already sorted
	for i := 0; i < n-1; i++ {
		sorted = true
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				sorted = false
			}
		}
		if sorted {
			return arr
		}
	}
	return arr
}
