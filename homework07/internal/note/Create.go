package note

import (
	"awesomeProject1/homework07/internal/models"
	"awesomeProject1/homework07/internal/utils"
	"awesomeProject1/homework07/internal/validators"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *NoteHandler) CreateNote(c *gin.Context) {
	userID, err := utils.GetUserID(c)
	if err != nil {
		utils.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	var req validators.CreateNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusUnprocessableEntity, "invalid note")
		return
	}

	note := models.Note{
		UserID:  userID,
		Title:   req.Title,
		Content: req.Content,
	}

	h.db.Create(&note)

	cacheKeyAllNotes := fmt.Sprintf("notes:user:%d", userID)
	h.cache.Del(c, cacheKeyAllNotes)

	utils.Success(c, note)
}
