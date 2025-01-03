package sort

import (
	"go-sort/constraints"
	"sync"
	"time"
)

func SleepSort[T constraints.Unsigned](arr []T) []T {
	var wg sync.WaitGroup
	ch := make(chan T, len(arr))
	for _, item := range arr {
		wg.Add(1)
		go func(item T) {
			defer wg.Done()
			time.Sleep(time.Duration(item) * time.Millisecond * 10)
			ch <- item
		}(item)
	}
	wg.Wait()
	close(ch)
	var result []T
	for item := range ch {
		result = append(result, item)
	}
	return result
}
