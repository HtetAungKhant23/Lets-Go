package main

import (
	"fmt"
	"testing"
)

func TestClaimPerTrip(t *testing.T) {
	type testCase struct {
		tripCount int
		expected  int
	}

	tests := []testCase{
		{30, 30000},
		{50, 50000},
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{31, 30500},
			{0, 0},
			{-1, 0},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		output := ClaimPerTrip(test.tripCount)
		if output != test.expected {
			failCount++
			t.Errorf(`------------------------------------
Inputs:     (%d)
Expected:   %d
Actual:     %d
Fail`, test.tripCount, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`------------------------------------
Inputs:     (%d)
Expected:   %d
Actual:     %d
Pass
`, test.tripCount, test.expected, output)
		}
	}
	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}
