package main

import (
	"fmt"
	"net/http"
	"simple-backend/api"
)

const port = ":8081"

func main() {
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	// Initialize Notes API
	notesAPI := api.NewNotes()

	// Setup HTTP handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!\n", s)
	})

	mux.HandleFunc("/api/v1/notes/list", notesAPI.GetNotes)
	mux.HandleFunc("/api/v1/notes/create", notesAPI.AddNote)
	mux.HandleFunc("/api/v1/notes/update", notesAPI.UpdateNote)
	mux.HandleFunc("/api/v1/notes/delete", notesAPI.DeleteNote)
	mux.HandleFunc("/api/v1/notes/find", notesAPI.GetNote)

	fmt.Println("Server running on port 8081...")
	if err := http.ListenAndServe(port, mux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
