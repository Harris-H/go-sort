package sort

import "go-sort/constraints"

func merge[T constraints.Ordered](a []T, b []T) []T {

	var r = make([]T, len(a)+len(b))
	var i = 0
	var j = 0

	for i < len(a) && j < len(b) {

		if a[i] <= b[j] {
			r[i+j] = a[i]
			i++
		} else {
			r[i+j] = b[j]
			j++
		}

	}

	for i < len(a) {
		r[i+j] = a[i]
		i++
	}
	for j < len(b) {
		r[i+j] = b[j]
		j++
	}

	return r

}

// Merge Perform merge sort on a slice

func MergeSort[T constraints.Ordered](items []T) []T {

	if len(items) < 2 {
		return items

	}

	var middle = len(items) / 2
	var a = MergeSort(items[:middle])
	var b = MergeSort(items[middle:])
	return merge(a, b)

}
