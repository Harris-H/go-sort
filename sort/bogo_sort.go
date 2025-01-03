package sort

import (
	"go-sort/constraints"
	"math/rand"
)

func isSorted[T constraints.Number](arr []T) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}

	return true
}

func shuffle[T constraints.Number](arr []T) {
	for i := 1; i < len(arr); i++ {
		j := rand.Intn(i)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func BogoSort[T constraints.Number](arr []T) []T {
	for !isSorted(arr) {
		shuffle(arr)
	}
	return arr
}
