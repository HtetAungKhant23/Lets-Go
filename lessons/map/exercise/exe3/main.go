package main

import (
	"fmt"
	"strings"
)

func getNameCounts(names []string) map[rune]map[string]int {
	res := make(map[rune]map[string]int)
	for _, name := range names {
		val, ok := res[rune(name[0])]
		if !ok {
			val = make(map[string]int)
			res[rune(name[0])] = val
		}
		val[name] += 1
	}
	return res
}

func countDistinctWords(messages []string) int {
	counter := 0
	counterMap := make(map[string]int)
	for _, msg := range messages {
		splitMsg := strings.Split(msg, " ")
		for _, m := range splitMsg {
			m = strings.ToLower(m)
			_, ok := counterMap[m]
			if !ok {
				counterMap[m] = 1
				counter++
			}
		}
	}
	return counter
}

func main() {
	//res := getNameCounts([]string{"😊", "htet", "koHAK", "khant", "htoo", "koHAK"})
	//fmt.Println(res)

	count := countDistinctWords([]string{
		"LF9M UBRS need all",
		"LF8M UBRS need all",
		"LF7M UBRS need all",
		"LF6M UBRS need tanks and heals",
		"LF5M UBRS need tanks and heals",
		"LF4M UBRS need tanks and heals",
		"LF3M UBRS need tanks and healer",
		"LF2M UBRS need tanks",
		"LF1M UBRS need tank",
		"Group is full thanks!",
	})
	fmt.Println(count)
}
