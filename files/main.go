package main

import (
	"fmt"
	"os"

	"github.com/htetaungkhant23/go/files/fileutils"
)

func main() {
	fmt.Println("Welcome from files.")
	rootPath, _ := os.Getwd()
	filePath := rootPath + "/data"
	c, err := fileutils.RaedTextFile(filePath + "/text.txt")
	if err != nil {
		fmt.Printf("Error found, %v", err)
	} else {
		fmt.Println(c)
		newContent := fmt.Sprintf("Original: %v\n New: %v %v", c, c, c)
		fileutils.WriteFile(filePath+"/output.txt", newContent)
	}
}
