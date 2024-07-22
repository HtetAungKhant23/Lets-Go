package main

import "fmt"

func init() {
	fmt.Println("Ohh.")
}

func init() {
	fmt.Println("nice.")
}

func calculate(price float32) (float32, float32) {
	return price * 0.5, price * 0.05
}

func changeAge(age *int) {
	*age++
}

func main() {
	// var tt float32 = 12.1
	// fmt.Println("Hello world")
	// one, _ := calculate(tt)
	// fmt.Println(one, tt)
	age := 23
	changeAge(&age)
	fmt.Println(age)
}
