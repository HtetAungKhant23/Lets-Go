package main

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		url      string
		expected []Issue
	}

	runCases := []testCase{
		{
			"https://api.boot.dev/v1/courses_rest_api/learn-http/issues?limit=1",
			[]Issue{{Title: "Fix that one bug nobody understands", Estimate: 19}},
		},
		{
			"https://api.boot.dev/v1/courses_rest_api/learn-http/issues?limit=2",
			[]Issue{
				{Title: "Fix that one bug nobody understands", Estimate: 19},
				{Title: "Implement user authentication flow", Estimate: 6},
			},
		},
	}

	for _, test := range runCases {
		issues, _ := getIssues(test.url)

		if !reflect.DeepEqual(issues, test.expected) {
			log.Fatalln("Fail")
		} else {
			fmt.Println("Pass")
		}
	}
}
