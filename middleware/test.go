package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func TestMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("=============")
		c.Next()
	}

}
