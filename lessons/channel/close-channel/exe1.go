package main

import (
	"fmt"
)

func countReports(numSentCh chan int) int {
	counter := 0
	for {
		reports, ok := <-numSentCh
		if !ok {
			break
		}
		counter += reports
	}
	return counter
}

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go sendReports(3, ch)
	reports := countReports(ch)

	fmt.Println("result reports", reports)
}
