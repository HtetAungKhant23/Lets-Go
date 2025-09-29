package main

import "fmt"

func sum(nums ...int) (total int) { // variadic function receives the variadic arguments as a slice.
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total
}

type cost struct {
	day   int
	value float64
}

func getCostsByDay(cost []cost) []float64 {
	var totalCost []float64
	for _, c := range cost {
		for len(totalCost) <= c.day {
			totalCost = append(totalCost, 0.0)
		}
		totalCost[c.day] += c.value
	}
	return totalCost
}

func main() {
	//nums := []int{}
	//total := sum(nums...) // spread operator allows us to pass a slice into a variadic function
	//fmt.Println(total)

	ex := []cost{
		{0, 4.0},
		{1, 2.1},
		{5, 2.5},
		{1, 3.1},
	}

	cos := getCostsByDay(ex)
	fmt.Println(cos)

}
