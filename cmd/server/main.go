package main

import (
	"fmt"
	"net/http"

	"github.com/alexnakagama/loca-go/internal/store"
	"github.com/alexnakagama/loca-go/internal/handler"
)

func main() {
	locationStore := store.NewStore()
	personStore := store.NewStore()

	locationHandler := handler.NewLocationHandler(locationStore)
	personHandler := handler.NewPersonHandler(personStore)

	http.HandleFunc("/locations", locationHandler.HandleLocation)
	http.HandleFunc("/persons", personHandler.HandlePerson)

	fmt.Println("server running on port: 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
