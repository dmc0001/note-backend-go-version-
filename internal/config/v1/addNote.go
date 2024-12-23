package config

type AddNoteRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
type AddNoteResponse struct {
	Result NoteResult
}

func (notes *Notes) AddNote(req *AddNoteRequest) AddNoteResponse {

	if req.Title == "" || req.Description == "" {
		return AddNoteResponse{
			Result: Failure{
				Message: "Title or description cannot be empty",
			},
		}
	}

	id := len(*notes) + 1

	note := Note{
		Id:          id,
		Title:       req.Title,
		Description: req.Description,
	}

	*notes = append(*notes, note)

	return AddNoteResponse{
		Result: Success{
			Note: note,
		},
	}

}
