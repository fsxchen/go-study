package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"blog/controllers/common"
	"blog/pkg/e"
	"blog/pkg/util"
)

func JWTHeadAuth(skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.Request.Header.Get("x-token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			c.Set(common.USER_ID_Key, claims.ID)
			c.Set(common.USER_NAME_Key, claims.Username)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}
		c.Next()
	}
}
