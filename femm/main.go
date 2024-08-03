package main

import (
	"fmt"
	"github.com/htetaungkhant/go/museum/api"
	"github.com/htetaungkhant/go/museum/data"
	"html/template"
	"net/http"
)

func getHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("templates/index.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error."))
		return
	}
	html.Execute(w, data.GetAll())
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello", getHello)
	server.HandleFunc("/template", handleTemplate)
	server.HandleFunc("/api/exhibition", api.Get)

	// using file server Handler to serve static file automatically
	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	err := http.ListenAndServe(":4004", server)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("server is running...")
}
