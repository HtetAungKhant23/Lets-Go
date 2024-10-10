package main

import "fmt"

type Analytic struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

func getMessageText(analytic *Analytic, msg Message) {
	(*analytic).MessagesTotal++ // can write in shorthand like -> analytic.MessagesTotal
	if msg.Success {
		(*analytic).MessagesSucceeded++
	} else {
		(*analytic).MessagesFailed++
	}
}

func main() {
	initialAnalytics := Analytic{MessagesTotal: 2, MessagesFailed: 1, MessagesSucceeded: 1}
	newMessage := Message{Recipient: "goofy", Success: false}
	fmt.Println("before ", initialAnalytics, newMessage)
	getMessageText(&initialAnalytics, newMessage)
	fmt.Println("after ", initialAnalytics, newMessage)
}
