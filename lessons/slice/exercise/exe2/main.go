package main

import "fmt"

type Message interface {
	Type() string
}

type TextMessage struct {
	Sender  string
	Context string
}

func (txtMsg *TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender    string
	MediaType string
	Context   string
}

func (mediaMsg MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender  string
	URL     string
	Context string
}

func (linkMsg LinkMessage) Type() string {
	return "link"
}

func filterMessage(messages []Message, filterType string) []Message {
	var res []Message
	for _, msg := range messages {
		if msg.Type() == filterType {
			res = append(res, msg)
			break
		}
	}
	return res
}

func main() {
	txtMsg := TextMessage{"hak", "asdfsasdf"}
	mediaMsg := MediaMessage{"susisa", "media", "asdfasdfasdfa"}
	res := filterMessage([]Message{&txtMsg, &mediaMsg}, "text")
	fmt.Println(res)

}
