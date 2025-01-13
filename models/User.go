package models

import "time"

type User struct {
	User_id    int
	Username   string
	Created_at time.Time
}
