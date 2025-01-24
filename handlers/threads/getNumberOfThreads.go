package threads

import (
	"context"
	"cvwo-winter-assignment/database"
	"strconv"
)

func GetNumberOfThreads(categoryStr string) (int, error) {
	countQuery := "SELECT COUNT(*) FROM threads"
	args := []interface{}{}

	if categoryStr != "" {
		categoryID, err := strconv.Atoi(categoryStr)
		if err != nil {
			return 0, err
		}
		countQuery += "WHERE category_id = $1"
		args = append(args, categoryID)
	}
	var totalThreads int
	err := database.Conn.QueryRow(context.Background(), countQuery).Scan(&totalThreads)

	if err != nil {
		return 0, err
	}
	return totalThreads, nil
}
