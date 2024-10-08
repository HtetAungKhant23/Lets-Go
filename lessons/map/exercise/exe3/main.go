package main

import "fmt"

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

func main() {
	res := getNameCounts([]string{"😊", "htet", "koHAK", "khant", "htoo", "koHAK"})
	fmt.Println(res)
}
