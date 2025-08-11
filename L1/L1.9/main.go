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
					println("conveer 1 is closed")
					close(ch2)
					return
				}
				println("put", a, "from ch1 to ch2")
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
					println("conveer 2 is closed")
					return
				}
				println("get", a, "from ch2")
			}
		}
	}()

	for i := 0; i < 10; i++ {
		println("put", i, "to ch1")
		ch1 <- i
	}
	close(ch1)
	wg.Wait()
	println("fabric is stopped")
}
