package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTagMessages(t *testing.T) {
	tests := []struct {
		messages []sms
		expected [][]string
	}{
		{
			messages: []sms{{id: "001", content: "Urgent, please respond!"}, {id: "002", content: "Big sale on all items!"}},
			expected: [][]string{{"Urgent"}, {"Promo"}},
		},
		{
			messages: []sms{{id: "003", content: "Enjoy your day"}},
			expected: [][]string{{}},
		},
		{
			messages: []sms{{id: "004", content: "Sale! Don't miss out on these urgent promotions!"}},
			expected: [][]string{{"Urgent", "Promo"}},
		},
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		actual := tagMessages(test.messages, tagger)
		if len(actual) != len(test.expected) {
			failCount++
			t.Errorf(`---------------------------------
Test Failed for length of returned sms slice
Expecting: %v
Actual:    %v
Fail`, len(test.expected), len(actual))
			continue
		}

		for i, msg := range actual {
			if !reflect.DeepEqual(msg.tags, test.expected[i]) {
				failCount++
				t.Errorf(`---------------------------------
Test Failed for message ID %s
Expecting: %v
Actual:    %v
Fail`, msg.id, test.expected[i], msg.tags)
			} else {
				passCount++
				fmt.Printf(`---------------------------------
Test Passed for message ID %s
Expecting: %v
Actual:    %v
Pass
`, msg.id, test.expected[i], msg.tags)
			}
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}