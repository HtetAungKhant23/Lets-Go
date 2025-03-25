package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getIssues(url string) ([]Issue, error) {
	var issue []Issue

	res, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err = json.Unmarshal(data, &issue); err != nil {
		return nil, fmt.Errorf("error decoding json: %w", err)
	}

	return issue, nil

}
