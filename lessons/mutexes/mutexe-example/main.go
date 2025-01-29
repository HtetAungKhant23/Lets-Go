package main

import (
	"fmt"
	"sync"
)

func writeLoop(m map[int]int, mu *sync.Mutex) {
	for {
		for i := 0; i < 100; i++ {
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}
	}
}

func readLoop(m map[int]int, mu *sync.Mutex) {
	for {
		mu.Lock()
		for k, v := range m {
			fmt.Println(k, '-', v)
		}
		mu.Unlock()
	}
}

func main() {
	m := map[int]int{}

	mu := &sync.Mutex{}

	go writeLoop(m, mu)
	go readLoop(m, mu)

	block := make(chan struct{})
	<-block
}
