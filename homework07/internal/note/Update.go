package note

import (
	"awesomeProject1/homework07/internal/models"
	"awesomeProject1/homework07/internal/utils"
	"awesomeProject1/homework07/internal/validators"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *NoteHandler) UpdateNote(c *gin.Context) {
	id := c.Param("id")
	userID, err := utils.GetUserID(c)
	if err != nil {
		utils.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	var req validators.UpdateNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}

	var note models.Note
	if err := h.db.Where("id = ? AND user_id = ?", id, userID).First(&note).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.Error(c, http.StatusNotFound, "note not found")
		} else {
			utils.Error(c, http.StatusInternalServerError, "database error")
		}
		return
	}

	update := make(map[string]interface{})
	if req.Title != nil {
		update["title"] = *req.Title
	}
	if req.Content != nil {
		update["content"] = *req.Content
	}

	if len(update) == 0 {
		utils.Success(c, note)
		return
	}

	if err := h.db.Model(&note).Updates(update).Error; err != nil {
		slog.Error("Update note failed", "error", err)
		utils.Error(c, http.StatusInternalServerError, "update failed")
		return
	}

	cacheKeyNote := "note:" + id
	cacheKeyAllNotes := fmt.Sprintf("notes:user:%d", userID)

	h.cache.Del(c, cacheKeyNote)
	h.cache.Del(c, cacheKeyAllNotes)

	h.db.First(&note, id)
	utils.Success(c, note)
}
