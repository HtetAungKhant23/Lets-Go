package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		rows, cols int
		expected   [][]int
	}

	testCases := []testCase{
		{3, 3, [][]int{
			{0, 0, 0},
			{0, 1, 2},
			{0, 2, 4},
		}},
		{4, 4, [][]int{
			{0, 0, 0, 0},
			{0, 1, 2, 3},
			{0, 2, 4, 6},
			{0, 3, 6, 9},
		}},
		{0, 0, [][]int{}},
		{5, 7, [][]int{
			{0, 0, 0, 0, 0, 0, 0},
			{0, 1, 2, 3, 4, 5, 6},
			{0, 2, 4, 6, 8, 10, 12},
			{0, 3, 6, 9, 12, 15, 18},
			{0, 4, 8, 12, 16, 20, 24},
		}},
	}

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		res := createMatrix(test.rows, test.cols)
		if !reflect.DeepEqual(res, test.expected) {
			failCount++
			fmt.Println("Fail")
		} else {
			passCount++
			fmt.Println("Pass")
		}
	}

	fmt.Println(passCount, failCount)
}
