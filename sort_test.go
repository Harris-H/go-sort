package main

import (
	"go-sort/sort"
	"math/rand"
	"testing"
	"time"
)

func generateRandomSlice(size int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	slice := make([]int, size)
	for i := range slice {
		slice[i] = r.Intn(10000) // generate random integers between 0 and 9999
	}
	return slice
}

func generateRandomSliceUint(size int) []uint {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	slice := make([]uint, size)
	for i := range slice {
		slice[i] = uint(r.Intn(10000)) // generate random integers between 0 and 9999
	}
	return slice
}

func BenchmarkBubble(b *testing.B) {
	// Adjust the size of the array as needed
	size := 1000
	arr := generateRandomSlice(size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Make a copy of the array to avoid sorting an already sorted array
		arrCopy := make([]int, size)
		copy(arrCopy, arr)
		sort.Bubble(arrCopy)
	}
}

func BenchmarkSelectSort(b *testing.B) {
	// Adjust the size of the array as needed
	size := 1000
	arr := generateRandomSlice(size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Make a copy of the array to avoid sorting an already sorted array
		arrCopy := make([]int, size)
		copy(arrCopy, arr)
		sort.SelectSort(arrCopy)
	}
}

func BenchmarkInsertSort(b *testing.B) {
	// Adjust the size of the array as needed
	size := 1000
	arr := generateRandomSlice(size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Make a copy of the array to avoid sorting an already sorted array
		arrCopy := make([]int, size)
		copy(arrCopy, arr)
		sort.InsertSort(arrCopy)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	// Adjust the size of the array as needed
	size := 1000
	arr := generateRandomSlice(size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Make a copy of the array to avoid sorting an already sorted array
		arrCopy := make([]int, size)
		copy(arrCopy, arr)
		sort.MergeSort(arrCopy)
	}
}

func BenchmarkShellSort(b *testing.B) {
	// Adjust the size of the array as needed
	size := 1000
	arr := generateRandomSlice(size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Make a copy of the array to avoid sorting an already sorted array
		arrCopy := make([]int, size)
		copy(arrCopy, arr)
		sort.ShellSort(arrCopy)
	}
}

func BenchmarkHeapSort(b *testing.B) {
	// Adjust the size of the array as needed
	size := 1000
	arr := generateRandomSlice(size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Make a copy of the array to avoid sorting an already sorted array
		arrCopy := make([]int, size)
		copy(arrCopy, arr)
		sort.HeapSort(arrCopy)
	}
}

func BenchmarkBucketSort(b *testing.B) {
	// Adjust the size of the array as needed
	size := 1000
	arr := generateRandomSliceUint(size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Make a copy of the array to avoid sorting an already sorted array
		arrCopy := make([]uint, size)
		copy(arrCopy, arr)
		sort.BucketSort(arrCopy)
	}
}

func BenchmarkRadixSort(b *testing.B) {
	// Adjust the size of the array as needed
	size := 1000
	arr := generateRandomSliceUint(size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Make a copy of the array to avoid sorting an already sorted array
		arrCopy := make([]uint, size)
		copy(arrCopy, arr)
		sort.RadixSort(arrCopy)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	// Adjust the size of the array as needed
	size := 1000
	arr := generateRandomSlice(size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Make a copy of the array to avoid sorting an already sorted array
		arrCopy := make([]int, size)
		copy(arrCopy, arr)
		sort.QuickSort(arrCopy)
	}
}

func BenchmarkBogoSort(b *testing.B) {
	// Adjust the size of the array as needed
	size := 10
	arr := generateRandomSlice(size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Make a copy of the array to avoid sorting an already sorted array
		arrCopy := make([]int, size)
		copy(arrCopy, arr)
		sort.BogoSort(arrCopy)
	}
}

//func BenchmarkSleepSort(b *testing.B) {
//	// Adjust the size of the array as needed
//	size := 10
//	arr := generateRandomSliceUint(size)
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		// Make a copy of the array to avoid sorting an already sorted array
//		arrCopy := make([]uint, size)
//		copy(arrCopy, arr)
//		sort.SleepSort(arrCopy)
//	}
//}
