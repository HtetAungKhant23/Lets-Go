package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSellProducts(t *testing.T) {
	type TestCase struct {
		books             []book
		toys              []toy
		expectedBookStore bookStore
		expectedToyStore  toyStore
	}

	tests := []TestCase{
		{
			books: []book{
				{
					title:  "Clean Code",
					author: "Robert C. Martin",
					price:  190.5,
				},
				{
					title:  "Titanic",
					author: "James Cameron",
					price:  120.4,
				},
			},
			toys: []toy{
				{
					name:  "Tom & Jelly",
					price: 35,
				},
				{
					name:  "Car",
					price: 26,
				},
			},
			expectedBookStore: bookStore{
				booksSold: []book{
					{
						title:  "Clean Code",
						author: "Robert C. Martin",
						price:  190.5,
					},
					{
						title:  "Titanic",
						author: "James Cameron",
						price:  120.4,
					},
				},
			},
			expectedToyStore: toyStore{
				toysSold: []toy{
					{
						name:  "Tom & Jelly",
						price: 35,
					},
					{
						name:  "Car",
						price: 26,
					},
				},
			},
		},
	}

	passCount := 0
	failCount := 0

	bs := bookStore{
		booksSold: []book{},
	}

	ts := toyStore{
		toysSold: []toy{},
	}

	for _, test := range tests {
		sellProducts[book](&bs, test.books)
		sellProducts[toy](&ts, test.toys)
		if reflect.DeepEqual(bs.booksSold, test.expectedBookStore.booksSold) && reflect.DeepEqual(ts.toysSold, test.expectedToyStore.toysSold) {
			passCount += 1
		} else {
			failCount += 1
		}
	}

	fmt.Println("Pass ", passCount)
	fmt.Println("Fail ", failCount)

}
