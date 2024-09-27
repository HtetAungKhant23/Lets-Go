package main

import "fmt"

type MemberShip struct {
	Type            string
	MemberCharLimit int
}

type User struct {
	Name string
	MemberShip
}

func newUser(name, membershipType string) User {
	charLimit := 100
	if membershipType == "premium" {
		charLimit = 1000
	}
	return User{
		Name: name,
		MemberShip: MemberShip{
			Type:            membershipType,
			MemberCharLimit: charLimit,
		},
	}
}

func (u *User) SendMessage(message string, messageLength int) (string, bool) {
	if u.MemberCharLimit < messageLength {
		return "", false
	}
	return message, true
}

func main() {
	user1 := newUser("hak", "normal")
	user2 := newUser("Phoo", "premium")
	//fmt.Println(user1, user2)

	fmt.Println(user1.SendMessage("wu shu war", 120))
	fmt.Println(user2.SendMessage("wu shu war", 120))
}
