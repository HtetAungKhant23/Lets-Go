package main

import "fmt"

func countConnections(groupSize int) (size int) {
	for i := 0; i < groupSize; i++ {
		size += groupSize - i - 1
	}
	return
}

func main() {
	size := countConnections(0)
	fmt.Println(size)
}
