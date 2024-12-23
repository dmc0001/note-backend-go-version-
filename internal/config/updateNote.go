package config

import "sort"

type UpdateNoteRequest struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateNoteResponse struct {
	Result NoteResult
}

func (notes *Notes) EditNote(req *UpdateNoteRequest) UpdateNoteResponse {

	index := sort.Search(len(*notes), func(i int) bool {
		return (*notes)[i].Id >= req.Id
	})

	//index := notes.FindNote(req.Id)

	if index < len(*notes) && (*notes)[index].Id == req.Id {
		(*notes)[index].Title = req.Title
		(*notes)[index].Description = req.Description
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

//func (notes *Notes) FindNote(id int) int {
//	for i, note := range *notes {
//		if note.Id == id {
//			return i
//		}
//	}
//	return -1
//}
