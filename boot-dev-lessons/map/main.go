package main

import "fmt"

type user struct {
	name        string
	phoneNumber int
}

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, fmt.Errorf("invalid sizes")
	}

	u := make(map[string]user)

	for i := 0; i < len(names); i++ {
		u[names[i]] = user{names[i], phoneNumbers[i]}
	}

	return u, nil
}

func main() {
	u, err := getUserMap([]string{"Eren", "Armin", "Mikasa"}, []int{14355550987, 98765550987, 18265554567})
	fmt.Println(u, err)
}
