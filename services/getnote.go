package services

import (
	"Ankipage/models"
	"errors"
)

func GetNote(noteID int) (*models.Note, error) {
	note, err := models.GetNoteByID(noteID)
	if err != nil {
		//c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return note, errors.New("error Note not found")
	}
	return note, nil
}
