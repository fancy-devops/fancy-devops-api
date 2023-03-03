package common

import (
	"net/http"

	"github.com/fancy-devops/fancy-devops-api/model/codes"
	"github.com/gin-gonic/gin"
)

type GinRes struct {
}

func NewGinRes() *GinRes {
	return &GinRes{}
}

func (obj *GinRes) ErrorResponse(c *gin.Context, code int) {
	httpCode := 200
	if code < 1000 {
		httpCode = code
	}

	c.JSON(httpCode, gin.H{
		"code": code,
		"msg":  codes.GetMsg(code),
		"data": nil,
	})
}

func (obj *GinRes) ErrorResponse2(c *gin.Context, code int, msg string) {
	httpCode := 200
	if code < 1000 {
		httpCode = code
	}

	c.JSON(httpCode, gin.H{
		"code": code,
		"msg":  codes.GetMsg(code) + ":" + msg,
		"data": nil,
	})
}

func (obj *GinRes) SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": codes.Common_Success,
		"msg":  codes.GetMsg(codes.Common_Success),
		"data": data,
	})
}
