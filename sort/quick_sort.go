package sort

import "go-sort/constraints"

func qSort[T constraints.Ordered](arr []T, l, r int) {
	if l < r {
		i, j, x := l, r, arr[l]
		for i < j {
			for ; i < j && arr[j] > x; j-- {
			}
			if i < j {
				arr[i] = arr[j]
				i++
			}
			for ; i < j && arr[i] < x; i++ {
			}
			if i < j {
				arr[j] = arr[i]
				j--
			}
		}
		arr[i] = x
		qSort(arr, l, i-1)
		qSort(arr, i+1, r)
	}
}
func QuickSort[T constraints.Ordered](arr []T) []T {
	qSort(arr, 0, len(arr)-1)
	return arr
}
