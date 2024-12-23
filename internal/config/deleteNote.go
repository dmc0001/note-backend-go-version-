package config

import "sort"

type DeleteNoteRequest struct {
	Id int `json:"id"`
}

type DeleteNoteResponse struct {
	Result NoteResult
}

func (notes *Notes) RemoveNote(req *DeleteNoteRequest) DeleteNoteResponse {

	if len(*notes) == 0 {
		return DeleteNoteResponse{
			Result: Failure{
				Message: "no notes available",
			},
		}
	}
	
	index := sort.Search(len(*notes), func(i int) bool {
		return (*notes)[i].Id >= req.Id
	})

	if index < len(*notes) && (*notes)[index].Id == req.Id {
		*notes = append((*notes)[:index], (*notes)[index+1:]...)
		return DeleteNoteResponse{
			Result: Success{(*notes)[index]},
		}
	}

	return DeleteNoteResponse{
		Result: Failure{
			Message: "note id does not exist",
		},
	}
}
