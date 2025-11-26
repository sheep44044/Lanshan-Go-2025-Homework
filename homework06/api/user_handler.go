package api

import (
	"awesomeProject1/homework06/dao"
	"awesomeProject1/homework06/model"
	"awesomeProject1/homework06/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
	}
	// 如果用户存在，这里这种是用户名可以一致的，即只要密码不一致就视为不同用户
	if dao.FindUser(req.Username, req.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user already exists",
		})
		return
	}
	dao.AddUser(req.Username, req.Password)
	c.JSON(http.StatusOK, gin.H{
		"message": "register success",
	})
}

func Login(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}
	// 检查用户是否存在且密码是否正确
	if !dao.FindUser(req.Username, req.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user not found",
		})
		return
	}
	// 生成jwt token
	token, err := utils.GenerateToken(req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}
	// 返回token
	c.JSON(http.StatusOK, gin.H{
		"message": "login",
		"token":   token,
		"refresh": refreshToken,
	})
}

func ModifyPassword(c *gin.Context) {
	// 从中间件中获取用户名
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "user not authenticated",
		})
		return
	}

	var req model.ModifyPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	// 验证旧密码是否正确
	if !dao.FindUser(username.(string), req.OldPassword) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "old password is incorrect",
		})
		return
	}

	dao.ModifyPassword(username.(string), req.NewPassword)
	c.JSON(http.StatusOK, gin.H{
		"message": "password changed successfully",
	})
}

func RefreshToken(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	// 验证refresh token
	token, err := utils.ValidateToken(req.RefreshToken)
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid refresh token",
		})
		return
	}

	claims, err := utils.ExtractClaims(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid token claims",
		})
		return
	}

	// 检查是否是refresh token
	if claims["type"] != "refresh" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not a refresh token",
		})
		return
	}

	username, ok := claims["username"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid token",
		})
		return
	}

	// 生成新的access token
	newToken, err := utils.GenerateToken(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	// 返回新的access token
	c.JSON(http.StatusOK, gin.H{
		"message": "token refreshed successfully",
		"token":   newToken,
	})
}
