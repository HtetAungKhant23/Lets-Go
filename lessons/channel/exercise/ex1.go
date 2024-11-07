package main

import "fmt"

func waitForDBs(numDBs int, dbChan chan struct{}) {
	fmt.Print(numDBs, dbChan)
	for i := 0; i < numDBs; i++ {
		<-dbChan
	}
}

func getDBsChannel(numDBs int) (chan struct{}, *int) {
	count := 0
	ch := make(chan struct{})

	fmt.Println(numDBs)

	go func() {
		for i := 0; i < numDBs; i++ {
			fmt.Println("hi")
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
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
