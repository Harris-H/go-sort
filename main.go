package main

import (
	"fmt"
	"go-sort/sort"
)

func main() {
	arr := []uint{6, 3, 2, 4, 5, 1, 8, 7, 12, 10}
	arr = sort.BogoSort(arr)
	fmt.Println(arr)
}
