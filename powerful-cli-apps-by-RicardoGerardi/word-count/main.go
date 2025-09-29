package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main(){
	fmt.Println(count(os.Stdin))
}

func count(r io.Reader) int {
	// A scanner is used to read text from a Reader (such as file)
	scanner := bufio.NewScanner(r)

	// Define the scanner split type to words (default is split by lines)
	scanner.Split(bufio.ScanWords)

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return  wc
}
