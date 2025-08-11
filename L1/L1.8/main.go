package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./main <N> <i>")
		return
	}

	n, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	i, err := strconv.ParseInt(os.Args[2], 10, 32)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if i < 1 || i > 9 {
		fmt.Println("Error: i must be between 1 and 9")
	}
	i--
	mask := int64(math.Pow(2, float64(i)))

	ans := n ^ mask // меняет i-ый бит на противоположный

	fmt.Println(ans)

}
