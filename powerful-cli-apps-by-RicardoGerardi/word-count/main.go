package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main(){
	// Defining a boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "Count lines")
	// Parsing the flags provided by the user
	flag.Parse()

	fmt.Println(count(os.Stdin, *lines))
}

func count(r io.Reader, countLines bool) int {
	// A scanner is used to read text from a Reader (such as file)
	scanner := bufio.NewScanner(r)

	if !countLines {
		// Define the scanner split type to words (default is split by lines)
		scanner.Split(bufio.ScanWords)
	}

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return  wc
}
