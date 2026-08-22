package main

import (
	"fmt"
	"net/http"

	"github.com/alexnakagama/loca-go/internal/store"
	"github.com/alexnakagama/loca-go/internal/handler"
)

func main() {
	store := store.NewStore()

	locationHandler := handler.NewLocationHandler(store)
	personHandler := handler.NewPersonHandler(store)

	http.HandleFunc("/locations", locationHandler.HandleLocation)
	http.HandleFunc("/persons", personHandler.HandlePerson)

	fmt.Println("server running on port: 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
