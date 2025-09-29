package main

import "fmt"

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

func updateBalance(cus *customer, transaction transaction) error {
	switch transaction.transactionType {
	case transactionDeposit:
		cus.balance += transaction.amount
	case transactionWithdrawal:
		if cus.balance < transaction.amount {
			return fmt.Errorf("insufficient funds")
		}
		cus.balance -= transaction.amount
	default:
		return fmt.Errorf("unknown transaction type")
	}
	return nil
}

func main() {
	alice := customer{id: 1, balance: 100.0}
	deposit := transaction{customerID: 1, amount: 50, transactionType: transactionDeposit}

	var _ error = updateBalance(&alice, deposit)
}
