package main

import (
	"fmt"
	"sync"
	"time"
)

type safeCounter struct {
	counts map[string]int
	mu     *sync.Mutex
}

func (sc safeCounter) inc(key string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.slowIncrement(key)
}

func (sc safeCounter) val(key string) int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.slowVal(key)
}

func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

func (sc safeCounter) slowVal(key string) int {
	time.Sleep(time.Microsecond)
	return sc.counts[key]
}

func main() {
	type testcase struct {
		email string
		count int
	}

	tests := []testcase{
		{"one@gmail.com", 23},
		{"two@gmail.com", 43},
	}

	for _, val := range tests {
		sc := safeCounter{
			counts: make(map[string]int),
			mu:     &sync.Mutex{},
		}
		var wg sync.WaitGroup
		for i := 0; i < val.count; i++ {
			wg.Add(1)
			go func(email string) {
				sc.inc(email)
				wg.Done()
			}(val.email)
		}
		wg.Wait()

		output := sc.val(val.email)
		fmt.Println(output, val.count)
	}
}
