// pointer receiver
// Methods with pointer receivers don't require that a pointer is used to call the method.
// The pointer will automatically be derived from the value.

package main

type email struct {
	message     string
	fromAddress string
	toAddress   string
}

func (e *email) setMessage(newMessage string) {
	e.message = newMessage
}

func main() {
	email := email{
		message:     "My name is Lt. Aldo Raine and I'm putting together a special team, and I need me eight soldiers.",
		fromAddress: "lt.aldo.raine@mailio.com",
		toAddress:   "army@mailio.com",
	}
	email.setMessage("You just say bingo.")
}
