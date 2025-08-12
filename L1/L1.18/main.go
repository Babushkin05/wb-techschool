package main

import "sync"

type AtomicCount struct {
	value int
	mutex sync.Mutex
}

func (c *AtomicCount) Increment() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value++
}

func (c *AtomicCount) Get() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.value
}

func main() {
	atomicCount := &AtomicCount{}
	wg := sync.WaitGroup{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			atomicCount.Increment()
		}()
	}

	wg.Wait()
	println(atomicCount.value)
}
