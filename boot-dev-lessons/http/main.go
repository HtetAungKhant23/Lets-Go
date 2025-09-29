package main

import (
	"fmt"
	"log"
)

const issueUrl = "https://api.boot.dev/v1/courses_rest_api/learn-http/issues"

func main() {
	issues, err := getIssueData(issueUrl)

	if err != nil {
		log.Fatalf("error getting issue data: %v", err)
	}

	fmt.Println(string(issues))

	prettyData, err := prettify(string(issues))

	if err != nil {
		log.Fatalf("error prettifying data: %v", prettyData)
	}

	fmt.Println(prettyData)

}
