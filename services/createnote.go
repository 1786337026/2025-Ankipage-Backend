package services

import (
	"Ankipage/models"
	"errors"
	"time"
)

func CreateNote(id int, userid int, content string, title string, url string) (int, error) {
	var note models.Note
	note.ID = id
	note.UserID = userid
	note.CreatedAt = time.Now()
	note.UpdatedAt = time.Now()
	note.Content = content
	note.Title = title
	note.Url = url
	models.Memory[userid] = append(models.Memory[userid], models.AnkiValue{note.ID, 2.5, 0, 0, time.Now(), 0})
	// 保存到数据库
	if err := note.Save(); err != nil {
		//c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save note"})
		return note.ID, errors.New("Failed to save note")
	}
	return note.ID, nil
}
