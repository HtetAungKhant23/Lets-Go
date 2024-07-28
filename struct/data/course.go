package data

import "fmt"

type Duration float32

type Course struct {
	Id         string
	Name       string
	Slug       string
	Legacy     bool
	Duration   Duration
	Instructor Instructor
}

func (c Course) Singup() bool {
	return true
}

func (c Course) String() string {
	return fmt.Sprintf("--- %v --- (%v)\n", c.Name, c.Instructor.FirstName)
}
