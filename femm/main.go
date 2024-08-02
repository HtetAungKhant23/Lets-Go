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

	// using file server Handler to serve static file automatically
	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	err := http.ListenAndServe(":4004", server)
	if err != nil {
		fmt.Println(err)
	}
}
