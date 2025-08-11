package main

import "sync"

func main() {
	m := sync.Map{}
	wg := sync.WaitGroup{}
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Store(i/2, i) // на каждый ключ претендуют 2 горутины
		}(i)
	}
	wg.Wait()
}

// go run L1.7/main.go -race
// ничего не выводит, потому что гонки не обнаружены
