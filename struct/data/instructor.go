package data

import "fmt"

type Instructor struct {
	Id        string
	FirstName string
	LastName  string
	Age       int
}

func NewInstructor(firstName string, lastName string) Instructor {
	return Instructor{FirstName: firstName, LastName: lastName}
}

func (i Instructor) PrettyPrint() string {
	return fmt.Sprintf("So prettry %v %v", i.FirstName, i.LastName)
}
