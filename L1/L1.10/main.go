package main

import "fmt"

func main() {
	arr := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	m := make(map[int][]float32)

	for _, v := range arr {
		a := 10 * (int(v) / 10)
		if _, ok := m[a]; !ok {
			m[a] = make([]float32, 0)
		}
		m[a] = append(m[a], v)
	}
	for k, v := range m {
		fmt.Printf("%d: %v\n", k, v)
	}
}
