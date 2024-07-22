package main

import (
	"fmt"
	"os"

	"github.com/htetaungkhant23/go/files/fileutils"
)

func main() {
	fmt.Println("Welcome from files.")
	rootPath, _ := os.Getwd()
	c, err := fileutils.RaedTextFile(rootPath + "/data/text.txt")
	if err != nil {
		fmt.Printf("Error found, ", err)
	} else {
		fmt.Println(c)
	}
}
