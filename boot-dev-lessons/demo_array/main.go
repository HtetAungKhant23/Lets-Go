package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

type Product struct {
	name  string
	price int
}

type authenticationInfo struct {
	username string
	password string
}

func checkProductInfo(prods [4]Product) {
	length := len(prods)
	totalCost := 0
	totalItems := 0
	for i := 0; i < length; i++ {
		if prods[i].name != "" {
			totalItems += 1
		}
		totalCost += prods[i].price
	}
	fmt.Println("total cost", totalCost)
	fmt.Println("total items", totalItems)
	fmt.Println("last items", prods[totalItems-1])
}

func checkCleanliness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		if rooms[i].cleaned {
			fmt.Println(rooms[i].name, "is clean.")
		} else {
			fmt.Println(rooms[i].name, "is dirty.")
		}
	}
}

func (auth authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %v:%v", auth.username, auth.password)
}

func main() {
	auth := authenticationInfo{"HAK", "HAK123"}
	fmt.Println(auth.getBasicAuth())

	//rooms := [...]Room{
	//	{"one", true},
	//	{"two", false},
	//	{"three", true},
	//	{"four", true},
	//}
	//checkCleanliness(rooms)

	//prods := [4]Product{
	//	{"one", 100},
	//	{"two", 200},
	//	{"three", 300},
	//}
	//
	//checkProductInfo(prods)
}
