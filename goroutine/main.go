package main

import (
	"fmt"
	"time"
)

func printText(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(800 * time.Millisecond)
	}
}

func main() {
	go printText("first")
	go printText("second")
	go printText("third")
	time.Sleep(time.Minute)
}
