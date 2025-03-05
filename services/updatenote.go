package services

import (
	"Ankipage/db"
	"Ankipage/models"
	"errors"
	"time"
)

func UpdateNote(id int, content string) error {
	var note models.Note

	// 查找原始笔记
	if err := db.DB.First(&note, id).Error; err != nil {
		//c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return errors.New("error Note not found")
	}

	// 更新笔记内容
	note.UpdatedAt = time.Now()
	note.Content = content
	if err := note.Update(); err != nil {
		//c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update note"})
		return errors.New("error Failed to Update Note")
	}
	return nil
}
