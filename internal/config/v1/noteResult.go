package v1

type Success struct {
	Note Note `json:"note"`
}

type Failure struct {
	Message string `json:"message"`
}

type NoteResult interface {
}

type Notes []Note
