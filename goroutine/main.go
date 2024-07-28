package main

import (
	"fmt"
	"time"
)

func printText(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(200 * time.Millisecond)
	}
	channel <- "Done"
}

func main() {
	// go printText("first")
	// go printText("second")
	// go printText("third")
	// time.Sleep(time.Minute)

	var channel chan string // channel := make(chan string)
	go printText("Passing channel with thread", channel)
	response := <-channel
	fmt.Println(response)
}
