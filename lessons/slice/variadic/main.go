package main

import "fmt"

func sum(nums ...int) (total int) { // variadic function receives the variadic arguments as a slice.
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total
}

func main() {
	nums := []int{}
	total := sum(nums...) // spread operator allows us to pass a slice into a variadic function
	fmt.Println(total)
}
