package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Request contins what the client sent
	// ResponseWriter is the object used to respond to the client
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Loca-Go API")
	})

	fmt.Println("server running on port: 8080")

	// This function tells to open a server in x port and start to accept requests
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
