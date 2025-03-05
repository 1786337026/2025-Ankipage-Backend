package services

import (
	"Ankipage/models"
	"errors"
)

func GetAllNote(UserID int) ([]models.Note, error) {

	note, err := models.GetAllNotes(UserID)

	if err != nil {
		//c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return []models.Note{}, errors.New("error Note not found")
	}

	return note, nil
}
