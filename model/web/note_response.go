package web

import "time"

type NoteResponse struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt bool `json:deleted_at`
}