package api

import (
	"encoding/json"
	"net/http"
	"simple-backend/internal/config/v1"
	"strconv"
)

type Notes struct {
	notes v1.Notes
}

func NewNotes() *Notes {
	return &Notes{
		notes: make(v1.Notes, 0),
	}
}

func (n *Notes) AddNote(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req v1.AddNoteRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response := n.notes.AddNote(&req)
	switch result := response.Result.(type) {

	case v1.Success:
		w.WriteHeader(http.StatusCreated)
		err := json.NewEncoder(w).Encode(result.Note)

		if err != nil {
			return
		}

	case v1.Failure:
		http.Error(w, result.Message, http.StatusBadRequest)
	}

}

func (n *Notes) UpdateNote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req v1.UpdateNoteRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	response := n.notes.EditNote(&req)
	switch result := response.Result.(type) {
	case v1.Success:
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(result.Note)
		if err != nil {
			return
		}

	case v1.Failure:
		http.Error(w, result.Message, http.StatusNotFound)
	}
}
func (n *Notes) DeleteNote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var req v1.DeleteNoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	response := n.notes.RemoveNote(&req)
	switch result := response.Result.(type) {

	case v1.Success:
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(result.Note)
		if err != nil {
			return
		}

	case v1.Failure:
		http.Error(w, result.Message, http.StatusNotFound)

	}

}

func (n *Notes) GetNotes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	if len(n.notes) == 0 {
		w.WriteHeader(http.StatusNoContent)
	}

	err := json.NewEncoder(w).Encode(n.notes)

	if err != nil {
		return
	}

}

func (n *Notes) GetNote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	idParam := r.URL.Query().Get("id")

	if idParam == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
		return
	}

	for _, note := range n.notes {
		if note.Id == id {
			// If found, respond with the note
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(note)
			if err != nil {
				http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			}
			return
		}
	}

	http.Error(w, "Note Not Found", http.StatusNotFound)

}
