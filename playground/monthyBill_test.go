package main

import (
	"fmt"
	"testing"
)

func TestMonthlyBill(t *testing.T) {
	type testCase struct {
		costPerSend  int
		numLastMonth int
		numThisMonth int
		expected     int
	}

	tests := []testCase{
		{costPerSend: 2, numLastMonth: 89, numThisMonth: 102, expected: 26},
		{costPerSend: 2, numLastMonth: 98, numThisMonth: 104, expected: 12},
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{3, 50, 40, -30},
			{3, 60, 60, 0},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		output := MonthlyBillIncrease(test.costPerSend, test.numLastMonth, test.numThisMonth)
		if output != test.expected {
			failCount++
			t.Errorf(`------------------------------------
Inputs:     (%d, %d, %d)
Expected:   %d
Actual:     %d
Fail`, test.costPerSend, test.numLastMonth, test.numThisMonth, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`------------------------------------
Inputs:     (%d, %d, %d)
Expected:   %d
Actual:     %d
Pass
`, test.costPerSend, test.numLastMonth, test.numThisMonth, test.expected, output)
		}
	}
	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}
