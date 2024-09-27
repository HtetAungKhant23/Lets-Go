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

func main() {
	c := circle{2}
	r := rect{12, 21}

	printShapeData(&c)
	printShapeData(r)

	bm := birthdayMessage{time.Now(), "hak"}
	sr := sendingReport{"Happy birthday", 3}

	printMessage(&bm)
	printMessage(&sr)

}
