package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT)
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	// launch workers
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				fmt.Println("worker", i, "exited")
				return
			}
		}(i)
	}

	<-sigs
	cancel()
	wg.Wait()
	fmt.Println("exiting")
}
