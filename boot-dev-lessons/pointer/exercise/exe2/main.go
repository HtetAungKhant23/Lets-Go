package main

import (
	"fmt"
	"strings"
)

// ------ you're dealing with pointers you should check if it's nil before trying to dereference it. ------

func removeProfanity(message *string) {
	if message == nil {
		return
	}
	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "fubb", "****")
	messageVal = strings.ReplaceAll(messageVal, "shiz", "****")
	messageVal = strings.ReplaceAll(messageVal, "witch", "*****")
	*message = messageVal
	fmt.Println("work")
}

func main() {
	msg := "English, motherfubber, do you speak it?"
	removeProfanity(nil)
	removeProfanity(&msg)
}
