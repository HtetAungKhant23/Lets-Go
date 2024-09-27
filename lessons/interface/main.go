package main

import (
	"fmt"
	"math"
	"time"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	height, width float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeData(s shape) {
	fmt.Println("Area =", s.area(), "and Perimeter =", s.perimeter())
}

type message interface {
	getMessage() string
}

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm *birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %v", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr *sendingReport) getMessage() string {
	return fmt.Sprintf("Your %s report is ready. You've sent %v messages.", sr.reportName, sr.numberOfSends)
}

func printMessage(m message) {
	fmt.Println(m.getMessage())
}

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c *contractor) getName() string {
	return c.name
}

func (c *contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

type fullTime struct {
	name   string
	salary int
}

func (ft *fullTime) getSalary() int {
	return ft.salary
}

func (ft *fullTime) getName() string {
	return ft.name
}

func printEmployeeData(e employee) {
	fmt.Println("Name", e.getName(), "And Salary", e.getSalary())
}

// ------- type implement multiple interface ----------

type Speaker interface {
	Speak() string
}

type Walker interface {
	Walk() string
}

type Robot struct {
	name string
}

func (r *Robot) Speak() string {
	return fmt.Sprintf("Beep beep! I am %s", r.name)
}

func (r *Robot) Walk() string {
	return fmt.Sprintf("%s can walk", r.name)
}

func printRobotData(s Speaker, w Walker) {
	fmt.Println(s.Speak())
	fmt.Println(w.Walk())

}

func main() {
	c := circle{2}
	r := rect{12, 21}

	printShapeData(&c) // cause of pointer receiver
	printShapeData(r)

	bm := birthdayMessage{time.Now(), "hak"}
	sr := sendingReport{"Happy birthday", 3}

	printMessage(&bm)
	printMessage(&sr)

	con := contractor{"Hak", 50, 2900}
	ft := fullTime{"Wu", 2000}

	printEmployeeData(&con)
	printEmployeeData(&ft)

	rob := Robot{"He He Robot"}
	printRobotData(&rob, &rob)

}
