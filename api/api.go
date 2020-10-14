package api

import (
	"brain-api/data"
	"brain-api/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Run is...
func Run() {
	r := mux.NewRouter()
	r.HandleFunc("/folders", InsertFolderHandler).Methods("POST")
	r.HandleFunc("/folders/{id}", UpdateFolderHandler).Methods("PUT")
	log.Println(http.ListenAndServe(":8000", r))
}

// InsertFolderHandler is...
func InsertFolderHandler(w http.ResponseWriter, r *http.Request) {
	var f models.Folder
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&f); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response, _ := json.Marshal("Invalid request")
		w.Write(response)
		return
	}
	defer r.Body.Close()

	f.Type = "folder"

	result, _ := data.InsertContentItem(&f)

	jsonResponse, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// UpdateFolderHandler is...
func UpdateFolderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var f models.Folder
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&f); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response, _ := json.Marshal("Invalid request")
		w.Write(response)
		return
	}
	defer r.Body.Close()

	f.Type = "folder"

	result, _ := data.UpdateContentItem(id, &f)

	jsonResponse, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
