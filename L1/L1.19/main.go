package main

func Reverse(s string) string {
	arr := []rune(s)
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}

	return string(arr)
}

func main() {
	str := "Hello, мир!💐屎"
	println(Reverse(str))
}
