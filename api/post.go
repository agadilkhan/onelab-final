package api

type PostRequest struct {
	Title   string `json:"title" db:"title"`
	Content string `json:"content" db:"content"`
}
