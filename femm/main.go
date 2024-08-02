package main

import (
	"fmt"
	"net/http"
)

func getHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":4004", server)
	if err != nil {
		fmt.Println(err)
	}
}
