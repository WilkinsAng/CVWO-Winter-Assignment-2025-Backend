package models

import "time"

type Thread struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	UserID     int       `json:"user_id"`
	Username   string    `json:"username"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Likes      int       `json:"likes"`
	Dislikes   int       `json:"dislikes"`
	CategoryID int       `json:"category_id"`
}
