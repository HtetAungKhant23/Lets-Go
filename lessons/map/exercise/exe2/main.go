// count instances

package main

func getCounts(messagedUser []string, validUsers map[string]int) {
	for _, u := range messagedUser {
		if _, ok := validUsers[u]; ok {
			validUsers[u]++
		}
	}
}

func main() {
	getCounts(
		[]string{"cersei", "jaime", "cersei"},
		map[string]int{"cersei": 0, "jaime": 0},
	)
}
