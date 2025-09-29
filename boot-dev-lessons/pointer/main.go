package main

import (
	"fmt"
	"strings"
)

// -------- you're dealing with pointers you should check if it's nil before trying to dereference it. ---------------

type Message struct {
	Recipient string
	Text      string
}

func getMessageText(m Message) string {
	return fmt.Sprintf(`
To: %v
Message: %v
`, m.Recipient, m.Text)
}

func removeProfanity(message *string) {
	replaceWord := map[string]string{
		"fubb":  "****",
		"shiz":  "****",
		"witch": "*****",
	}
	fmt.Println("before ", message, *message)
	for k, word := range replaceWord {
		*message = strings.ReplaceAll(*message, k, word)
	}
	fmt.Println("after ", message, *message)
}

func main() {
	s := "Oh man I've seen some crazy ass shiz in my time..."
	removeProfanity(&s)

	msg := Message{"asdf", "jk;lk"}
	fmt.Println(getMessageText(msg))
}
