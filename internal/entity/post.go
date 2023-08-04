package entity

type Post struct {
	ID int64 `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
	Content string `json:"content" db:"content"`
	UserID int64 `json:"user_id" db:"user_id"`
}