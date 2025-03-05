package services

import (
	"Ankipage/models"
	"errors"
	"time"
)

func AnkiPush(ID int, Memory int, userID int) error {
	for index, note := range models.Memory[userID] {
		if note.NoteId == ID {
			models.Memory[userID][index].Anki(Memory)
			models.Memory[userID][index].Date = models.Memory[userID][index].Date.Add(time.Duration(models.Memory[userID][index].Value) * time.Hour)
			//c.JSON(http.StatusOK)
			return nil
		}
	}
	return errors.New("error there is no this note")
}
