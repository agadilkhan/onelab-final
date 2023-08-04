package api

type CreatePostRequest struct {
	Title string `json:"title" db:"title"`
	Content string `json:"content" db:"content"`
}