// message tagger exercise

//1. Complete the tagMessages function. It should take a slice of sms messages, and a function (that takes a sms as input and returns a slice of strings) as inputs.
//And it should return a slice of sms messages.
//It should loop through each message and set the tags to the result of the passed in function.
//Be sure to modify the messages of the original slice.
//Ensure the tags field contains an initialized slice. No nil slices.

//2. Complete the tagger function. It should take a sms message and return a slice of strings.
//For any message that contains "urgent" or "Urgent" in the content, the Urgent tag should be applied first.
//For any message that contains "sale" or "Sale" in the content, the Promo tag should be applied second.

package main

import (
	"fmt"
	"strings"
)

type sms struct {
	id      string
	content string
	tags    []string
}

func main() {
	messages := []sms{
		{id: "001", content: "Urgent! Last chance to see!"},
		{id: "002", content: "Big sale on all items!"},
		// Additional messages...
	}
	taggedMessages := tagMessages(messages, tagger)
	fmt.Println(taggedMessages)
	// `taggedMessages` will now have tags based on the content.
	// 001 = [Urgent]
	// 002 = [Promo]
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	for index, msg := range messages {
		tags := tagger(msg)
		messages[index].tags = tags
	}
	return messages
}

func tagger(msg sms) []string {
	tags := []string{}
	if strings.Contains(msg.content, "Urgent") || strings.Contains(msg.content, "urgent") {
		tags = append(tags, "Urgent")
	}
	if strings.Contains(msg.content, "Sale") || strings.Contains(msg.content, "sale") {
		tags = append(tags, "Promo")
	}

	return tags
}
