package main

import "fmt"

func concurrentFib(n int) []int {
	var numSlice []int
	ch := make(chan int)

	go fibonacci(n, ch)
	for value := range ch {
		numSlice = append(numSlice, value)
	}

	return numSlice
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	num := 5

	numSlice := concurrentFib(num)

	fmt.Println("number slice", numSlice)
}
