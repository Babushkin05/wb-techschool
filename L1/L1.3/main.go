package main

import (
	"fmt"
	"os"
	"strconv"
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
	for i := 0; i < n; i++ {
		go func(i int) {
			for {
				fmt.Printf("worker %v read %v\n", i, <-ch)
			}
		}(i)
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	i := 0
	for {
		select {
		case <-ticker.C:
			ch <- i
			i++
		}
	}
}
