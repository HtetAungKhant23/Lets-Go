package main

import "fmt"

func main() {
	fmt.Println("Calculator")
	fmt.Println("----------")
	var operation string
	var num1, num2 int
	fmt.Println("Enter operation! (add, sub, mul, div)")
	fmt.Scanf("%s", &operation)
	fmt.Println("Enter first number")
	fmt.Scanf("%d", &num1)
	fmt.Println("Enter secont number")
	fmt.Scanf("%d", &num2)
	fmt.Print("The answer is ")
	switch operation {
	case "add":
		fmt.Println(num1 + num2)
	case "sub":
		fmt.Println(num1 - num2)
	case "mul":
		fmt.Println(num1 * num2)
	case "div":
		fmt.Println(num1 / num2)
	}
}
