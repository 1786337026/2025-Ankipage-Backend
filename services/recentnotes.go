package services

import (
	"Ankipage/models"
	"errors"
)

func RecentNotes(userid int, limit int) ([]models.Note, error) {
	notes, err := models.GetRecentNotes(userid, limit)
	if err != nil {
		//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notes"})
		return []models.Note{}, errors.New("error Failed to fetch notes")
	}
	return notes, nil
}
