package main

import (
	"fmt"
	"testing"
)

func TestCountGroupConnection(t *testing.T) {
	type testCase struct {
		groupSize int
		expected  int
	}

	testCases := []testCase{
		{1, 0},
		{2, 1},
		{3, 3},
		{4, 6},
		{5, 10},
		{0, 0},
		{10, 45},
		{100, 4950},
		{1000, 499500},
	}

	passCount := 0
	failCount := 0

	for _, tc := range testCases {
		res := countConnections(tc.groupSize)
		if res != tc.expected {
			failCount++
			t.Errorf(`---------------------------------
Group Size: %v
Expecting: %v
Actual:    %v
Fail`, tc.groupSize, tc.expected, res)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Group Size: %v
Expecting: %v
Actual:    %v
Pass
`, tc.groupSize, tc.expected, res)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}
