package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}

	sort.Ints(a)
	sort.Ints(b)

	ans := make([]int, 0)
	l := 0
	r := 0

	for l < len(a) && r < len(b) {
		if a[l] < b[r] {
			l++
		} else if a[l] > b[r] {
			r++
		} else {
			ans = append(ans, a[l])
			l++
		}
	}

	fmt.Printf("%v", ans)
}
