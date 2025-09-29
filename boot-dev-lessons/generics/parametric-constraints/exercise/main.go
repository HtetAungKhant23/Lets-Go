package main

import (
	"fmt"
)

type biller[C customer] interface {
	Charge(C) bill
	Name() string
}

type customer interface {
	GetBillingEmail() string
}

type bill struct {
	Customer customer
	Amount   float64
}

type user struct {
	UserEmail string
}

func (u user) GetBillingEmail() string {
	return u.UserEmail
}

type userBiller struct {
	Plan string
}

func (ub userBiller) Charge(u user) bill {
	amount := 50.0
	if ub.Plan == "pro" {
		amount = 100.0
	}
	return bill{
		Customer: u,
		Amount:   amount,
	}
}

func (ub userBiller) Name() string {
	return fmt.Sprintf("%s user biller plan", ub.Plan)
}

type org struct {
	Admin user
	Name  string
}

func (o org) GetBillingEmail() string {
	return o.Admin.GetBillingEmail()
}

type orgBiller struct {
	Plan string
}

func (ob orgBiller) Charge(o org) bill {
	amount := 2000.0
	if ob.Plan == "pro" {
		amount = 3000.0
	}
	return bill{
		Customer: o,
		Amount:   amount,
	}
}

func (ob orgBiller) Name() string {
	return fmt.Sprintf("%s org biller plan", ob.Plan)
}

func billerFunc[C customer](b biller[C], c C) string {
	currentBill := b.Charge(c)
	return fmt.Sprintf("%f bill of %s", currentBill.Amount, currentBill.Customer.GetBillingEmail())
}

func main() {
	userBills := billerFunc[user](userBiller{Plan: "basic"}, user{UserEmail: "htet@gmail.com"})
	orgBills := billerFunc[org](orgBiller{Plan: "pro"}, org{Admin: user{UserEmail: "aung@gmail.com"}, Name: "Aung Khant"})
	fmt.Println(userBills)
	fmt.Println(orgBills)
}
