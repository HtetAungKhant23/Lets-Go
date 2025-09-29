package main

import (
	"fmt"
	"log"
)

const URL = "https://api.boot.dev/v1/courses_rest_api/learn-http/issues?limit=1"

func main() {
	data, err := getIssues(URL)

	if err != nil {
		log.Fatalf("error getting issue: %v", err)
	}

	fmt.Println(data)
}
