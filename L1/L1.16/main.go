package main

import (
	"fmt"
	"math/rand"
)

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	} else {
		pivot := arr[0]
		var less []int
		var greater []int
		for _, value := range arr[1:] {
			if value <= pivot {
				less = append(less, value)
			} else {
				greater = append(greater, value)
			}
		}
		less = QuickSort(less)
		greater = QuickSort(greater)
		arr = append(less, pivot)
		arr = append(arr, greater...)
		return arr
	}
}

func main() {
	arr := make([]int, 20)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}

	fmt.Println("Случайный массив:", arr)

	arr = QuickSort(arr)

	fmt.Println("Отсортированный массив:", arr)
}
