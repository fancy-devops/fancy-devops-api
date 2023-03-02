package jwt

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.chad122.top/fancy-devops/fancy-devops-api/model/codes"
	"gitlab.chad122.top/fancy-devops/fancy-devops-api/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int

		code = codes.Common_Success
		token := c.GetHeader("token")
		if token == "" {
			code = codes.Error_User_AuthTokenNotExist
		} else {
			claims, err := util.NewJwt().ParseToken(token)
			if err != nil {
				code = codes.Error_User_AuthTokenParseError
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = codes.Error_User_AuthTokenExpired
			}
		}

		if code != codes.Common_Success {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  codes.GetMsg(code),
				"data": "",
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
