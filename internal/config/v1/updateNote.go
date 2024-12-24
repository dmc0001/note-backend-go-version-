package v1

import "sort"

type UpdateNoteRequest struct {
	Id          int    `json:"id"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type UpdateNoteResponse struct {
	Result NoteResult
}

func (notes *Notes) EditNote(req *UpdateNoteRequest) UpdateNoteResponse {

	index := sort.Search(len(*notes), func(i int) bool {
		return (*notes)[i].Id >= req.Id
	})

	if index < len(*notes) && (*notes)[index].Id == req.Id {

		if req.Title != "" {
			(*notes)[index].Title = req.Title
		}

		if req.Description != "" {
			(*notes)[index].Description = req.Description
		}

		return UpdateNoteResponse{
			Result: Success{(*notes)[index]},
		}
	}

	return UpdateNoteResponse{
		Result: Failure{
			Message: "note id does not exist",
		},
	}
}
