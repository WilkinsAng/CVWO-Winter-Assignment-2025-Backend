package middleware

import (
	"context"
	"cvwo-winter-assignment/database"
	"errors"
)

var (
	ErrCommentNotFound     = errors.New("Comment not found.")
	ErrUnauthorizedComment = errors.New("You are unauthorized to update/delete this comment.")
)

/*
Check is user is the owner of the comment, used in update and delete
*/
func ValidateCommentOwnership(commentID int, userID int) error {

	var commentOwnerID int

	userQuery := `SELECT user_id FROM comments WHERE id = $1`

	err := database.Conn.QueryRow(context.Background(), userQuery, commentID).Scan(&commentOwnerID)

	if err != nil {
		return ErrCommentNotFound
	}

	if commentOwnerID != userID {
		return ErrUnauthorizedComment
	}

	return nil
}
