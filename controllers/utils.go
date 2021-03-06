package controllers

import (
	"github.com/gin-gonic/gin"
	"elephant-api/common/errCode"
)

func Response(c *gin.Context, httpCode, code int, data interface{}) {
	if httpCode == 500 {
		c.Status(500)
	}

	c.JSON(httpCode, gin.H{
		"msg":  errCode.ErrorString(code),
		"code": code,
		"data": data,
	})
}
