package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a handler function for handling requests
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, Go Web Server!")
	}

	// Register the handler function for a specific route pattern ("/")
	http.HandleFunc("/", handler)

	// Start the server on port 8080
	port := 8080
	fmt.Printf("Server is listening on :%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
