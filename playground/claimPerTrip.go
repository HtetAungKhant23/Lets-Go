package main

var mmkPerTrip int = 500

func claimPerTrip(trip int) (totalAmount int) {
	if trip >= 80 {
		totalAmount = calcMMK(80, trip)
	} else if trip >= 50 {
		totalAmount = calcMMK(50, trip)
	} else if trip >= 30 {
		totalAmount = calcMMK(30, trip)
	} else {
		totalAmount = trip * mmkPerTrip
	}
	return totalAmount
}

func calcMMK(limit int, trip int) (total int) {
	total = ((trip - limit) * mmkPerTrip) + (limit * 1000)
	return total
}
