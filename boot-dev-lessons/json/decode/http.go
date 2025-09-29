package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getIssue(url string) ([]Issue, error) {
	var issue []Issue
	res, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("error creating request %w", err)
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	if err = decoder.Decode(&issue); err != nil {
		return nil, fmt.Errorf("error decoding response data %w", err)
	}

	return issue, nil
}
