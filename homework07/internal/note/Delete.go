package note

import (
	"awesomeProject1/homework07/internal/models"
	"awesomeProject1/homework07/internal/utils"
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

	result := h.db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Note{})
	if result.RowsAffected == 0 {
		utils.Error(c, http.StatusNotFound, "note not found or permission denied")
		return
	}

	slog.Info("Cache cleared for deleted note", "note_id", id)
	utils.Success(c, gin.H{"message": "deleted"})
}
