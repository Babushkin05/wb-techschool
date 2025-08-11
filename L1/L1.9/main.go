package main

import "sync"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case a, ok := <-ch1:
				if !ok {
					close(ch2)
					return
				}
				ch2 <- a * 2
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case a, ok := <-ch2:
				if !ok {
					return
				}
				println(a)
			}
		}
	}()

	for i := 0; i < 10; i++ {
		ch1 <- i
	}
	close(ch1)
	wg.Wait()
}
