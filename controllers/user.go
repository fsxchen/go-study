package controllers

import (
	"blog/dto"
	"blog/service"

	"github.com/gin-gonic/gin"
)

var userService = service.UserService{}

type UserController struct {
}

func (u *UserController) List(c *gin.Context) {
	var listDto dto.GeneralListDto

	data, total := userService.List(listDto)

	resp(c, map[string]interface{}{
		"result": data,
		"total":  total,
	})
}

func (u *UserController) Create(c *gin.Context) {
	var userDto dto.UserCreateDto

	user, err := userService.Create(userDto)

	if err != nil {
		fail(c, ErrInputData)
		return
	}

	resp(c, map[string]interface{}{
		"result": user,
	})
}

// func AccountCtl(c *gin.Context) {
// 	var user models.User
// 	accounts, err := user.ListAccount()
// 	if err == nil {
// 		fmt.Println(accounts)
// 	} else {
// 		fmt.Println(err)
// 	}
// 	// name := c.Param("name")
// 	c.JSON(http.StatusOK, gin.H{
// 		"code": 0,
// 		"data": accounts,
// 	})
// }

// // @Summary Add a new pet to the store
// // @Description get string by ID
// // @Accept  json
// // @Produce  json
// // @Param   name     body    string     true        "name"
// // @Param   email     body    string     true        "name"
// // @Success 200 {string} string	"ok"
// // @Router /string [post]
// func RegisterCtl(c *gin.Context) {
// 	type RequestBody struct {
// 		Username string `json:"username" binding:"required"`
// 		Email    string `json:"email" binding:"required"`
// 	}

// 	var body RequestBody

// 	if err := c.BindJSON(&body); err != nil {
// 		c.AbortWithStatus(409)
// 		return
// 	}

// 	user := models.User{
// 		Name:  body.Username,
// 		Email: body.Email,
// 	}

// 	user.Save()

// 	c.JSON(201, gin.H{
// 		"code": 0,
// 	})

// }
