package services

import (
	"Ankipage/models"
	"errors"
	"time"
)

func Anki(userid int) ([]models.Note, error, bool) {
	var notes []models.Note
	flag := false
	for _, note := range models.Memory[userid] {
		if note.Date.Before(time.Now()) {
			flag = true
			//models.Memory[UserID][index].Anki(memory)
			ans, err := models.GetNoteByID(note.NoteId)
			if err != nil {
				//c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
				return notes, errors.New("error Note not found"), flag
			}
			notes = append(notes, *ans)
			//models.Memory[UserID][index].Date = models.Memory[UserID][index].Date.Add(time.Duration(models.Memory[UserID][index].Value) * time.Hour)
			//c.JSON(http.StatusOK, ans)
		}
	}
	return notes, nil, flag
}
