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
	r.HandleFunc("/folders", InsertFolderHandler).Methods("POST")
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

	newID := data.InsertOneContentItem(&f)
	fmt.Println(newID)

	jsonResponse, _ := json.Marshal(f)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
