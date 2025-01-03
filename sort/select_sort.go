package sort

import "go-sort/constraints"

func SelectSort[T constraints.Ordered](arr []T) []T {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		k := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[k] {
				k = j
			}
		}
		arr[i], arr[k] = arr[k], arr[i]
	}
	return arr
}
