package main

import "fmt"

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	costForCus, err := sendSMS(msgToCustomer)
	costForSpouse, e := sendSMS(msgToSpouse)
	if err != nil || e != nil {
		return 0, fmt.Errorf("%v", err)
	}
	return costForCus + costForSpouse, nil
}

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}

func main() {
	totalCost, err := sendSMSToCouple("Good Shal", "Shal Good kwar")
	fmt.Println(totalCost, err)
}
