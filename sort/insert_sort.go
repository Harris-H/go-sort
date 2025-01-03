// description: Implementation of insertion sort algorithm
// worst-case time complexity: O(n^2)
// average-case time complexity: O(n^2)
// space complexity: O(1)
// stable: yes

package sort

import "go-sort/constraints"

func InsertSort[T constraints.Ordered](arr []T) []T {
	for i := 1; i < len(arr); i++ {
		x := arr[i]
		j := i
		for ; j > 0 && arr[j-1] > x; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = x
	}
	return arr
}
