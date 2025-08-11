package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	ans := make([]string, 0)
	seen := make(map[string]struct{})
	for _, v := range arr {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			ans = append(ans, v)
		}
	}

	fmt.Printf("%v\n", ans)
}
