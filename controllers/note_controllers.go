package controllers

import (
	"Ankipage/models"
	"Ankipage/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Create a new note
// @Description Create a note for a user
// @Tags Notes
// @Accept json
// @Produce json
// @Param userid path int true "User ID"
// @Param note body models.Note true "Note Data"
// @Success 200 {object} map[string]interface{} "Note created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 500 {object} map[string]interface{} "Failed to save note"
// @Router /notes/{userid} [post]
func CreateNote(c *gin.Context) {
	userID := c.Param("userid")
	UserID, _ := strconv.Atoi(userID)
	var note models.Note
	if err := c.BindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":                           1,
			"error: Invalid request body   ": err.Error(),
		})
		return
	}

	noteId, err := services.CreateNote(c.GetInt("id"), UserID, note.Content, note.Title, note.Url)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  1,
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "Note created successfully",
		"data": gin.H{
			"noteID": noteId,
		},
	})
}

// @Summary List recent notes
// @Description Get the most recent notes for a user
// @Tags Notes
// @Accept json
// @Produce json
// @Param userid path int true "User ID"
// @Success 200 {array} models.Note "List of recent notes"
// @Failure 500 {object} map[string]interface{} "Failed to fetch notes"
// @Router /notes/recent/{userid} [get]
func ListRecentNotes(c *gin.Context) {
	userID := c.GetInt("userid")
	limit := 4

	// 获取最近的笔记

	notes, err := services.RecentNotes(userID, limit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  1,
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "Recent notes successfully",
		"data":    notes,
	})
}

// @Summary Get a specific note
// @Description Get a note by its ID
// @Tags Notes
// @Accept json
// @Produce json
// @Param id path int true "Note ID"
// @Success 200 {object} models.Note "Note data"
// @Failure 404 {object} map[string]interface{} "Note not found"
// @Router /notes/{id} [get]
func GetNote(c *gin.Context) {
	id := c.Param("id")
	noteID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":  1,
			"error": "Invalid note id",
		})
	}
	// 获取笔记

	note, err := services.GetNote(noteID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "Get Note successfully",
		"data":    note,
	})
}

// @Summary Get all notes
// @Description Get all notes for a user
// @Tags Notes
// @Accept json
// @Produce json
// @Param userid path int true "User ID"
// @Success 200 {array} models.Note "List of all notes"
// @Failure 404 {object} map[string]interface{} "Notes not found"
// @Router /notes/all/{userid} [get]
func GetNotes(c *gin.Context) {
	userID := c.Param("userid")
	UserID, _ := strconv.Atoi(userID)
	fmt.Println(UserID)
	// 获取笔记

	note, err := services.GetAllNote(UserID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":  1,
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "Get Notes successfully",
		"data":    note,
	})
}

// @Summary Update a note
// @Description Update a specific note by its ID
// @Tags Notes
// @Accept json
// @Produce json
// @Param id path int true "Note ID"
// @Param note body models.Note true "Updated Note Data"
// @Success 200 {object} map[string]interface{} "Note updated successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 404 {object} map[string]interface{} "Note not found"
// @Failure 500 {object} map[string]interface{} "Failed to update note"
// @Router /notes/{id} [put]
func UpdateNote(c *gin.Context) {
	id := c.Param("id")
	Id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"code":  1,
			"error": "Invalid ID",
		})
		return
	}
	var note models.Note
	if err := c.BindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  1,
			"error": "Invalid request body",
		})
		return
	}
	// 查找原始笔记

	err = services.UpdateNote(Id, note.Content)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  1,
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "Note updated successfully",
	})
}

// @Summary Delete a note
// @Description Delete a specific note by its ID
// @Tags Notes
// @Accept json
// @Produce json
// @Param id path int true "Note ID"
// @Success 200 {object} map[string]interface{} "Note deleted successfully"
// @Failure 400 {object} map[string]interface{} "Invalid ID"
// @Failure 500 {object} map[string]interface{} "Failed to delete note"
// @Router /notes/{id} [delete]
func DeleteNote(c *gin.Context) {
	id := c.Param("id")
	noteID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"code":  1,
			"error": "Invalid ID",
		})
		return
	}
	// 删除笔记

	err = services.DeleteNote(noteID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  1,
			"error": "Failed to delete note",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "Note deleted successfully",
	})
}
