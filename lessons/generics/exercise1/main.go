package main

import "fmt"

func getLast[T any](s []T) T {
	length := len(s)
	if length == 0 {
		var zero T
		return zero
	}
	return s[length-1]
}

func main() {
	res := getLast([]int{1, 2, 3, 4})
	fmt.Println(res)
}
