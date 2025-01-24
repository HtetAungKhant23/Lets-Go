package main

import "fmt"

func addEmailsToQueue(emails []string) chan string {
	ch := make(chan string, len(emails))
	go func() {
		for i := 0; i < len(emails); i++ {
			ch <- emails[i]
		}
	}()
	return ch
}

func sendEmails(size int, ch chan string) {
	for i := 0; i < size; i++ {
		email := <-ch
		fmt.Println("sending to", email)
	}
}

func main() {
	emails := []string{"htet@gmail.com", "htetaung@gmail.com", "htetaungkhant@gmail.com"}
	ch := addEmailsToQueue(emails)
	sendEmails(len(emails), ch)
}
