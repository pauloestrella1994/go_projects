package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(response, "404 not found", http.StatusNotFound)
		return
	}
	if request.Method != "GET" {
		http.Error(response, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprint(response, "hello!")
}

func formHandler(response http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		fmt.Fprintf(response, "ParseForm() err: %v\n", err)
		return
	}
	fmt.Fprint(response, "POST request successful\n")
	name := request.FormValue("name")
	address := request.FormValue("address")
	fmt.Fprintf(response, "Name = %v\n", name)
	fmt.Fprintf(response, "Address = %v", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
