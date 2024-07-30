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
	fmt.Println("trying to give")
	channel <- "Done"
}

func main() {
	// go printText("first")
	// go printText("second")
	// go printText("third")
	// time.Sleep(time.Minute)

	// var channel chan string // just declare can be nil
	channel := make(chan string) // declare and initialize
	fmt.Println(channel)
	go printText("Passing channel with thread", channel)
	response := <-channel
	fmt.Println(response)
}
