package main

import "fmt"

func waitForDBs(numDBs int, dbChan chan struct{}) {
	for i := 0; i < numDBs; i++ {
		<-dbChan
	}
}

func getDBsChannel(numDBs int) (chan struct{}, *int) {
	count := 0
	ch := make(chan struct{})

	go func() {
		for i := 0; i < numDBs; i++ {
			ch <- struct{}{}
			count++
		}
	}()

	return ch, &count
}

func main() {
	type testCase struct {
		numDbs int
	}

	t := testCase{3}

	ch, count := getDBsChannel(t.numDbs)
	waitForDBs(t.numDbs, ch)
	fmt.Println(ch, *count)

}
