// package main

// import "fmt"

// func waitForDBs(numDBs int, dbChan chan struct{}) {
// 	fmt.Println("in wait for dbs ", numDBs, dbChan)
// 	for i := 0; i < numDBs; i++ {
// 		fmt.Println("loop in wait for dbs", i)
// 		<-dbChan
// 	}
// }

// func getDBsChannel(numDBs int) (chan struct{}, *int) {
// 	fmt.Println("in get dbs channel")
// 	count := 0
// 	ch := make(chan struct{})
// 	fmt.Println("channel is made", ch)

// 	go func() {
// 		for i := 0; i < numDBs; i++ {
// 			fmt.Println("before put data in channel")
// 			ch <- struct{}{}
// 			fmt.Printf("Database %v is online\n", i+1)
// 			count++
// 		}
// 	}()

// 	return ch, &count
// }

// func main() {
// 	type testCase struct {
// 		numDbs int
// 	}

// 	t := testCase{3}

// 	ch, count := getDBsChannel(t.numDbs)
// 	fmt.Println("result from get db channel", *count, ch)
// 	waitForDBs(t.numDbs, ch)
// 	fmt.Println(ch, *count)

// }
