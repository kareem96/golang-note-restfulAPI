package web
type NoteUpdateRequest struct {
	Id int `validate:"required"`
	Title string `validate:"required,max=100,min=1" json:"title"`
	Description string `validate:"required,max=200,min=1" json:"description"`
}