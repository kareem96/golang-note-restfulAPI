package web

type NoteResponse struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}