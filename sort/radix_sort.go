package sort

import "go-sort/constraints"

func RadixSort[T constraints.Unsigned](arr []T) []T {
	var i, d, x T
	i, d, x = 0, 1, 10
	n := len(arr)
	ans := make([]T, n)
	for i := 0; i < n; i++ {
		if arr[i] >= x {
			x *= 10
			d++
		}
	}
	for i, x = 0, 1; i < d; i++ {
		w := make([]T, 10)
		for j := 0; j < n; j++ {
			w[(arr[j]/x)%10]++
		}
		for j := 1; j < 10; j++ {
			w[j] += w[j-1]
		}
		for j := n - 1; j >= 0; j-- {
			p := (arr[j] / x) % 10
			ans[w[p]-1] = arr[j]
			w[p]--
		}
		ans, arr = arr, ans
		x *= 10
	}
	return arr
}
