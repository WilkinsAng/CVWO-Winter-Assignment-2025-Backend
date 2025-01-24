package categories

import (
	"context"
	"cvwo-winter-assignment/database"
	"cvwo-winter-assignment/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCategories(c *gin.Context) {
	query := `SELECT id, name FROM categories ORDER BY name ASC`
	rows, err := database.Conn.Query(context.Background(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}
	defer rows.Close()

	var categories []models.Categories

	for rows.Next() {
		var category models.Categories
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning categories"})
			return
		}
		categories = append(categories, category)
	}

	c.JSON(http.StatusOK, gin.H{"categories": categories})
}
