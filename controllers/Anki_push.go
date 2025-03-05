package controllers

import (
	"Ankipage/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary      Get Anki Notes
// @Description  Get Anki Notes for user to read.
// @Tags         Notes
// @Accept       json
// @Produce      json
// @Param        userid  path      int     true   "User ID"
// @Param        memory  path      int     true   "Memory level (Anki score)"
// @Success      200     {object}  models.Note  "Success response with the updated note"
// @Failure      404     {object}  map[string]interface{}  "Error response if note not found"
// @Router       /anki/{userid}/ [get]
func Anki(c *gin.Context) {
	userID := c.Param("userid")
	UserID, _ := strconv.Atoi(userID)
	//temp := c.Param("memory")
	//memory, _ := strconv.Atoi(temp)

	notes, err, flag := services.Anki(UserID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if !flag {
		c.JSON(http.StatusOK, gin.H{"Congratulations!": "there is no note needed to read"})
	} else {
		c.JSON(http.StatusOK, gin.H{"Notes": notes})
	}
}

func AnkiPush(c *gin.Context) {
	temp := c.Param("id")
	ID, _ := strconv.Atoi(temp)
	temp = c.Param("memory")
	Memory, _ := strconv.Atoi(temp)
	temp = c.Param("userid")
	userID, _ := strconv.Atoi(temp)

	err := services.AnkiPush(ID, userID, Memory)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":  1,
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
	})
}
