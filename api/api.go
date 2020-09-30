package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Run is...
func Run() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/hello/{there}", HelloHandler).Methods("GET")
	log.Println(http.ListenAndServe(":8000", r))
}

// HomeHandler is...
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	response := Test{Message: "Hello World"}
	jsonResponse, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// HelloHandler is...
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	response := Test{Message: "Hello " + vars["there"]}
	jsonResponse, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// Test is...
type Test struct {
	Message string
}
