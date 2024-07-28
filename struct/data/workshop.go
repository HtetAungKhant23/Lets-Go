package data

import "time"

type Workshop struct {
	Course     // embedding course -> can have all method and variable of that struct
	Instructor // embedding instructor
	date       time.Time
}

func (w Workshop) Singup() bool {
	return true
}

func NewWorkshop(name string, instructor Instructor) Workshop {
	w := Workshop{}
	w.Name = name
	w.Instructor = instructor
	return w
}
