package main

import "encoding/json"

// The json.Marshal function converts a Go struct into a slice of bytes representing JSON data.

func marshalAll[T any](items []T) ([][]byte, error) {
	var res [][]byte

	for _, item := range items {
		if data, err := json.Marshal(item); err != nil {
			return nil, err
		} else {
			res = append(res, data)
		}
	}

	return res, nil
}
