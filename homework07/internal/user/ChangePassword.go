package user

import (
	"awesomeProject1/homework07/internal/models"
	"awesomeProject1/homework07/internal/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (h *UserHandler) ModifyPassword(c *gin.Context) {
	username := c.GetString("username")
	if username == "" {
		utils.Error(c, http.StatusUnauthorized, "user not authenticated")
		return
	}

	var req models.PasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, "invalid request")
		return
	}

	var user models.User
	if err := h.db.Select("id, password").Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.Error(c, http.StatusUnauthorized, "user not found")
		} else {
			utils.Error(c, http.StatusInternalServerError, "database error")
		}
		return
	}

	newHash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "failed to hash new password")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)); err != nil {
		utils.Error(c, http.StatusUnauthorized, "old password is incorrect")
		return
	}

	if err := h.db.Model(&user).Update("password", string(newHash)).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "failed to update password")
		return
	}

	utils.Success(c, gin.H{"message": "password changed successfully"})
}
