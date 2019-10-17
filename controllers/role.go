package controllers

import (
	"blog/schema"

	"github.com/gin-gonic/gin"
)

func AddRole(c *gin.Context) {
	var s_role schema.Role

	if err := c.BindJSON(&s_role); err != nil {
		c.AbortWithStatus(409)
		return
	}

	// s_role.Save()

	c.JSON(201, gin.H{
		"code": 0,
	})
}
