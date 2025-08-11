package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	// condition
	wg.Add(1)
	go func(a int) {
		defer wg.Done()
		time.Sleep(time.Millisecond * 10)
		for {
			if a == 1 {
				fmt.Println("condition")
				return
			}
		}
	}(1)

	// channel
	ticker := time.Tick(time.Millisecond * 10)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ticker:
				fmt.Println("chanel")
				return
			}
		}
	}()

	// context
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
	defer cancel()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("context")
				return
			}
		}
	}()

	// runtime.Goexit()
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 10)
		fmt.Println("Goexit")
		runtime.Goexit()
	}()

	wg.Wait()

	// // panic
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	panic("panic for all program")
	// }()

	// wg.Wait()

}
