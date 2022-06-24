package main

import (
	"fmt"
	"net/http"
	"go-server-basic/pkg/handlers"

)

const portNumber = ":8080"

// main is the main application fuction
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)

}
