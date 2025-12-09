package note

import (
	"awesomeProject1/homework07/internal/models"
	"awesomeProject1/homework07/internal/utils"
	"awesomeProject1/homework07/internal/validators"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *NoteHandler) UpdateNote(c *gin.Context) {
	id := c.Param("id")
	userid, exists := c.Get("user_id")
	if !exists {
		utils.Error(c, http.StatusUnauthorized, "未登录")
		return
	}

	userIDStr, ok := userid.(string)
	if !ok {
		utils.Error(c, http.StatusInternalServerError, "用户ID类型错误")
		return
	}
	// 将字符串转回 uint
	uid, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "用户ID格式错误")
		return
	}
	userID := uint(uid)

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

	h.db.First(&note, id)
	utils.Success(c, note)
}
