package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(response http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Println(response, "ParseForm() err: %v", err)
		return
	}

	firstName := request.FormValue("first_name")
	lastName := request.FormValue("last_name")

	fmt.Fprint(response, "Message succesfully sent \n")

	fmt.Fprint(response, "First Name: ", firstName)
	fmt.Fprint(response, "\n Last Name: ", lastName)
}

func helloHandler(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(response, "Page requested is not founded", http.StatusNotFound)
		return
	}

	if request.Method != "GET" {
		http.Error(response, "This method is not allowed", http.StatusMethodNotAllowed)
	}

	fmt.Fprint(response, "Hello page here")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting is running on 8080 port \n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
