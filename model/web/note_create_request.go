package web
type NoteCreateRequest struct {
	Title string `validate:"required,min=1,max=100" json:"title"`
	Description string `validate:"required,min=1,max=100" json:"description"`
	UserId int64 `validate:"required" json:"user_id"`
}