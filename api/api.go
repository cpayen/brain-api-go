package api

import (
	"brain-api/data"
	"brain-api/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Run is...
func Run() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/hello/{there}", HelloHandler).Methods("GET")
	r.HandleFunc("/insert/folder", InsertFolderHandler).Methods("GET")
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

// InsertFolderHandler is...
func InsertFolderHandler(w http.ResponseWriter, r *http.Request) {
	folder := models.NewFolder("coucou")
	newID := data.InsertOneFolder(folder)
	fmt.Println(newID)
	jsonResponse, _ := json.Marshal(folder)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
