package middleware

import (
	"context"
	"cvwo-winter-assignment/database"
	"errors"
)

var (
	ErrThreadNotFound = errors.New("Thread not found.")
	ErrUnauthorized   = errors.New("You are unauthorized to update/delete this thread.")
)

/*
Check is user is the owner of the thread, used in update and delete
*/
func ValidateThreadOwnership(threadID int, userID int) error {

	var threadOwnerID int

	userQuery := `SELECT user_id FROM threads WHERE id = $1`

	err := database.Conn.QueryRow(context.Background(), userQuery, threadID).Scan(&threadOwnerID)

	if err != nil {
		return ErrThreadNotFound
	}

	if threadOwnerID != userID {
		return ErrUnauthorized
	}

	return nil
}
