package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}

	i := 2

	copy(a[i:], a[i+1:])
	a = a[:len(a)-1]

	fmt.Printf("a: %v\n", a)
}
