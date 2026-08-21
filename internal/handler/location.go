package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alexnakagama/loca-go/interal/model"
)

func HandleLocation(w http.ResponseWriter, r *http.Request) {
	var location model.Location

	err := json.NewDecoder(r.Body).Decode(&location)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Println(location)
}
