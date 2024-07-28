package main

import (
	"fmt"

	data "github.com/htetaungkhant/go/server/data"
)

func main() {
	instructor := data.Instructor{Id: "1", FirstName: "Htet Aung", LastName: "Khant", Age: 25}
	println(instructor.PrettyPrint())

	instructor2 := data.NewInstructor("Kyle", "Simpson")

	course1 := data.Course{"1", "GoLang", "https://frontendmaster/go", true, 1.5, instructor2}
	fmt.Printf("%v\n\n", course1)

	w := data.NewWorkshop("oh", instructor2)
	// fmt.Printf("\n%v", w)

	courses := [2]data.Singable{}
	courses[0] = course1
	courses[1] = w

	for i, course := range courses {
		fmt.Println(i, course)
	}

}
