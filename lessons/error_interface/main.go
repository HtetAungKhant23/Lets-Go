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

func getSMSErrorString(cost float64, recipient string) string {
	return fmt.Sprintf("SMS that costs %v to be sent to %s cannot be sent", cost, recipient)
}

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf(getSMSErrorString(float64(costPerChar*len(message)), "09123456789"))
	}
	return costPerChar * len(message), nil
}

// ------------------ custom error ------------------

type divideError struct {
	dividend float64
}

// custom error by implementing built-in Error interface
func (e *divideError) Error() string {
	return fmt.Sprintf("can not divide %v by zero", e.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, &divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}

func main() {
	//totalCost, err := sendSMSToCouple("Good Shal asdfa asdfqe2 aswer asdzcarg asdfrqsf", "Shal Good kwar")
	//fmt.Println(totalCost, err)

	d := divideError{8}
	res, err := divide(d.dividend, 0)
	fmt.Println(res, err)
}
