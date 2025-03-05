package services

import (
	"Ankipage/models"
	"errors"
)

func DeleteNote(noteID int) error {
	err := models.DeleteNoteByID(noteID)
	if err != nil {
		return errors.New("error Failed to delete note")
	}
	return nil
}
