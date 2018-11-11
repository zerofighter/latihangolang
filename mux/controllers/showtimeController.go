package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"../common"
	"../data"
)

// Handler for HTTP Get - "/showtimes"
// Returns all Showtime documents
func GetShowTimes(w http.ResponseWriter, r *http.Request) {
	// Get all showtimes form repository
	showtimes := data.GetAll()
	j, err := json.Marshal(GoTestResource{Data: showtimes})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

var GetShowTimeById =  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// Get date from incoming url
	vars := mux.Vars(r)
	id  := vars["id"]
	// Create new context
	showtimes,err := data.GetAllbyID(id)

	if err != nil{
		json.Marshal(nil)
	}
	j, err := json.Marshal(GoTestResourceOne{Data: showtimes})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
})
