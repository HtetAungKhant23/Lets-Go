package main

import "fmt"

func indexOfFirstBadWord(messages []string, badWords []string) int {
	for i, msg := range messages {
		for _, bad := range badWords {
			if msg == bad {
				return i
			}
		}
	}
	return -1
}

func main() {
	str1 := []string{"crap", "shoot", "frick", "dang"}
	str2 := []string{""}

	idx := indexOfFirstBadWord(str1, str2)
	fmt.Println(idx)
}
