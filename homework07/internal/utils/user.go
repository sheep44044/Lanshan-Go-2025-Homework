package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func GetUserID(c *gin.Context) (uint, error) {
	uidRaw, exists := c.Get("user_id")
	if !exists {
		return 0, errors.New("未登录")
	}

	if uidFloat, ok := uidRaw.(float64); ok {
		return uint(uidFloat), nil
	}

	if uidUint, ok := uidRaw.(uint); ok {
		return uidUint, nil
	}

	return 0, errors.New("用户ID格式错误")
}
