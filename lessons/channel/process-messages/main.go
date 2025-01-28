package main

import (
	"fmt"
	"time"
)

func processMessages(messages []string) []string {
	var processedMessages []string
	ch := make(chan string)

	go func() {
		for _, message := range messages {
			ch <- process(message)
		}
		close(ch)
	}()

	for v := range ch {
		processedMessages = append(processedMessages, v)
	}

	return processedMessages
}

func process(message string) string {
	time.Sleep(1 * time.Second)
	return message + "-processed"
}

func main() {
	messages := []string{"Here are some messages", "Here is another", "and another"}
	processedMessages := processMessages(messages)
	fmt.Println(processedMessages)
}
