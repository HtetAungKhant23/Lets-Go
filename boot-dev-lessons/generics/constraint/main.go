package main

import (
	"errors"
	"fmt"
	"time"
)

type subscription struct {
	userEmail string
	startDate time.Time
	interval  string
}

type oneTimeUsagePlan struct {
	userEmail        string
	numEmailsAllowed int
}

// custom constraint for using in generic
type lineItem interface {
	GetCost() float64
	GetName() string
}

func chargeForLineItem[T lineItem](newItem T, oldItems []T, balance float64) ([]T, float64, error) {
	cost := newItem.GetCost()
	if cost > balance {
		var zeroVal []T
		return zeroVal, float64(0), errors.New("insufficient funds")
	}
	oldItems = append(oldItems, newItem)
	balance -= cost
	return oldItems, balance, errors.New("")
}

func (s subscription) GetName() string {
	return fmt.Sprintf("%s subscription", s.interval)
}

func (s subscription) GetCost() float64 {
	if s.interval == "monthly" {
		return 25.00
	}
	if s.interval == "yearly" {
		return 250.00
	}
	return 0.0
}

func (otup oneTimeUsagePlan) GetName() string {
	return fmt.Sprintf("one time usage plan with %v emails", otup.numEmailsAllowed)
}

func (otup oneTimeUsagePlan) GetCost() float64 {
	const costPerEmail = 0.03
	return float64(otup.numEmailsAllowed) * costPerEmail
}

func main() {
	newItem := lineItem(subscription{
		userEmail: "hak@gmail.com",
		startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		interval:  "yearly",
	})

	oldItems := []lineItem{
		subscription{
			userEmail: "oknasa@gmail.com",
			startDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			interval:  "monthly",
		},
		oneTimeUsagePlan{
			userEmail:        "triss@maribor",
			numEmailsAllowed: 100,
		},
	}

	userBalance := 1000.00

	items, updBalance, error := chargeForLineItem(newItem, oldItems, userBalance)

	fmt.Println(items, updBalance, error)
}
