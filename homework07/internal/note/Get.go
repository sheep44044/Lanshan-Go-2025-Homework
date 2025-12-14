package note

import (
	"awesomeProject1/homework07/internal/models"
	"awesomeProject1/homework07/internal/utils"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *NoteHandler) GetNote(c *gin.Context) {
	userID, err := utils.GetUserID(c)
	if err != nil {
		utils.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	id := c.Param("id")
	cacheKey := "note:" + id

	cachedNote, err := h.cache.Get(c, cacheKey)
	if err == nil {
		var note models.Note
		if err := json.Unmarshal([]byte(cachedNote), &note); err == nil {
			slog.Debug("Notes retrieved from cache", "key", cacheKey)
			utils.Success(c, note)
			return
		}
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

	noteJSON, _ := json.Marshal(note)
	h.cache.SetWithRandomTTL(c, cacheKey, string(noteJSON), 10*time.Minute)

	utils.Success(c, note)
}

func (h *NoteHandler) GetNotes(c *gin.Context) {
	userID, err := utils.GetUserID(c)
	if err != nil {
		utils.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	cacheKey := fmt.Sprintf("notes:user:%d", userID)
	cachedNotes, err := h.cache.Get(c, cacheKey)
	if err == nil {
		var notes []models.Note
		if err := json.Unmarshal([]byte(cachedNotes), &notes); err == nil {
			slog.Debug("Notes retrieved from cache", "key", cacheKey)
			utils.Success(c, notes)
			return
		}
	}

	var notes []models.Note
	err = h.db.Where("user_id = ?", userID).
		Order("updated_at DESC").
		Find(&notes).Error

	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "database error")
		return
	}

	notesJSON, _ := json.Marshal(notes)
	h.cache.SetWithRandomTTL(c, cacheKey, string(notesJSON), 10*time.Minute)

	utils.Success(c, notes)
}
