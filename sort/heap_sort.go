package sort

import "go-sort/constraints"

// HeapSort 堆排序
func heapifyDown[T constraints.Ordered](arr []T, n, i int) {
	j := (i << 1) + 1
	for ; j < n; i, j = j, (i<<1)+1 {
		if j+1 < n && arr[j] < arr[j+1] {
			j++
		}
		if arr[i] < arr[j] {
			arr[i], arr[j] = arr[j], arr[i]
		} else {
			break
		}
	}
}
func HeapSort[T constraints.Ordered](arr []T) []T {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		heapifyDown(arr, n, i)
	}
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapifyDown(arr, i, 0)
	}
	return arr
}
