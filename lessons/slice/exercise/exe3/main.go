package main

import "fmt"

func isValidPassword(password string) bool {
	existUpperCase := false
	existDigit := false
	if len(password) < 5 || len(password) > 12 {
		return false
	}
	for _, character := range password {
		switch {
		case int(character) >= 65 && int(character) <= 90:
			existUpperCase = true
		case int(character) >= 48 && int(character) <= 57:
			existDigit = true
		}
		if existUpperCase && existDigit {
			return true
		}
	}
	return false
}

func main() {
	res := isValidPassword("Password1")
	fmt.Println(res)
}
