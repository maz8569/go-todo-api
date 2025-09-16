package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/tasks", listHandler).Methods("GET")
	r.HandleFunc("/tasks", addHandler).Methods("POST")
	r.HandleFunc("/tasks/{id:[0-9]+}/done", doneHandler).Methods("PUT")
	r.HandleFunc("/tasks/{id:[0-9]+}", deleteHandler).Methods("DELETE")

	fmt.Println("Server is running on http://localhost:8080")

	http.ListenAndServe(":8080", r)
}
