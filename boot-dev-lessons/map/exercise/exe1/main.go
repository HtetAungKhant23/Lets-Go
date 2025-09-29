package main

import "fmt"

type user struct {
	name                string
	number              int
	scheduleForDeletion bool
}

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	u, ok := users[name]
	if !ok {
		return false, fmt.Errorf("not found")
	}

	if !u.scheduleForDeletion {
		return false, nil
	}

	delete(users, name)
	return true, nil
}

func main() {
	res, err := deleteIfNecessary(map[string]user{"Erwin": {"Erwin", 14355550987, true}, "Levi": {"Levi", 98765550987, true}, "Hanji": {"Hanji", 18265554567, true}}, "Erwin")
	fmt.Println(res, err)
}
