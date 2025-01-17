package models

import "time"

type User struct {
	User_id    int       `json:"user_id"`
	Username   string    `json:"username"`
	Created_at time.Time `json:"created_at"`
}
