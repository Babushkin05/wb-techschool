package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./main <N>")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	ch := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < n; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	after := time.After(time.Second * time.Duration(n))
	for {
		select {
		case <-after:
			fmt.Println("stopping")
			wg.Wait()
			return
		case i := <-ch:
			fmt.Printf("readed from channel: %d\n", i)
		}
	}

}
