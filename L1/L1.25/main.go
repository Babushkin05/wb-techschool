package main

import "time"

func Sleep(t time.Duration) {
	ticker := time.NewTicker(t)
	<-ticker.C
}

func main() {
	println("start")
	Sleep(time.Second)
	println("finish")
}
