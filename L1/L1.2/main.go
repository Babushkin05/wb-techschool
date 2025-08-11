package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}

	ans := 0

	wg := sync.WaitGroup{}
	for _, v := range arr {
		wg.Add(1)
		go func() {
			ans += v * v
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(ans)
}
