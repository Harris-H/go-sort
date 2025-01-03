package sort

import "go-sort/constraints"

// ShellSort 希尔排序
func ShellSort[T constraints.Ordered](arr []T) []T {
	n := len(arr)
	for d := n / 2; d > 0; d >>= 1 {
		for i := d; i < n; i++ {
			x, j := arr[i], i
			for ; j >= d && arr[j-d] > x; j -= d {
				arr[j] = arr[j-d]
			}
			arr[j] = x
		}
	}
	return arr
}
