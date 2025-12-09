package note

import (
	"awesomeProject1/homework07/internal/models"
	"awesomeProject1/homework07/internal/utils"
	"awesomeProject1/homework07/internal/validators"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *NoteHandler) CreateNote(c *gin.Context) {
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

	utils.Success(c, note)
}
