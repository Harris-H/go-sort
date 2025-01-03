package sort

import (
	"go-sort/constraints"
)

// Only Support Unsigned Integer
// BucketSort 桶排序

func BucketSort[T constraints.Unsigned](arr []T) []T {
	n := len(arr)
	mx := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] > mx {
			mx = arr[i]
		}
	}
	b := make([]int, mx+1)
	ans := make([]T, n)
	var i T
	for i := 0; i < n; i++ {
		b[arr[i]]++
	}
	for i = 1; i <= mx; i++ {
		b[i] += b[i-1]
	}
	for i := n - 1; i >= 0; i-- {
		ans[b[arr[i]]-1] = arr[i]
		b[arr[i]]--
	}
	return ans
}
