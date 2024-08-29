package main

func GetMonthlyPrice(tier string) int {
	switch tier {
	case "basic":
		return changeDollarToPennie(100)
	case "premium":
		return changeDollarToPennie(150)
	case "enterprise":
		return changeDollarToPennie(500)
	default:
		return 0
	}
}

func changeDollarToPennie(dollar int) int {
	return dollar * 100
}
