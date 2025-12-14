package note

import (
	"awesomeProject1/homework07/internal/models"
	"awesomeProject1/homework07/internal/utils"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *NoteHandler) DeleteNote(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id <= 0 {
		utils.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	userID, err := utils.GetUserID(c)
	if err != nil {
		utils.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	result := h.db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Note{})
	if result.RowsAffected == 0 {
		utils.Error(c, http.StatusNotFound, "note not found or permission denied")
		return
	}

	cacheKeyNote := "note:" + c.Param("id")
	cacheKeyAllNotes := fmt.Sprintf("notes:user:%d", userID)

	h.cache.Del(c, cacheKeyNote)
	h.cache.Del(c, cacheKeyAllNotes)

	slog.Info("Cache cleared for deleted note", "note_id", id)
	utils.Success(c, gin.H{"message": "deleted"})
}
