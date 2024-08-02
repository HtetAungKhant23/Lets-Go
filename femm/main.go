package main

import (
	"fmt"
	"net/http"
)

func getHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":4004", nil)
	if err != nil {
		fmt.Println(err)
	}
}
