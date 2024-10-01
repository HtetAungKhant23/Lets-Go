package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		nums     []int
		expected int
	}

	testCases := []testCase{
		{[]int{1, 2, 3, 4}, 10},
		{[]int{5, 6, 1, 2}, 14},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 120},
		{[]int{}, 0},
		{[]int{5}, 5},
	}

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		res := sum(test.nums...)
		if res != test.expected {
			failCount++
			t.Errorf("Fail")
		} else {
			passCount++
			fmt.Println("Pass")
		}
	}

	fmt.Println("Pass ", passCount, "fail", failCount)
}
