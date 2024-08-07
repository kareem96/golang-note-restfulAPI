package web
type NoteCreateRequest struct {
	Id int `validate:"required" json:"id"`
	Title string `validate:"required,min=1,max=100" json:"title"`
	Description string `validate:"required,min=1,max=100" json:"description"`
}