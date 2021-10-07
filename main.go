package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("restSanity Started")
	r := mux.NewRouter()
	r.HandleFunc("/", Test_Get).Methods("GET")
	r.HandleFunc("/", Test_Post).Methods("POST")
	log.Fatal(http.ListenAndServe(":5001", r))
	log.Println("restSanity Ended")
}

func Test_Get(w http.ResponseWriter, r *http.Request) {
	log.Println("Test_Get")
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Get: %v\n", vars["test"])
}

func Test_Post(w http.ResponseWriter, r *http.Request) {
	log.Println("Test_Post")
	var json_data map[string]string
	if err := json.NewDecoder(r.Body).Decode(&json_data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Printf("Posted: %v\n", json_data)
	fmt.Fprintf(w, "Posted: %v\n", json_data)
}
