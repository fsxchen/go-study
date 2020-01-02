package v1

import (
	"blog/controllers/common"
	"blog/model"
	"blog/pkg/e"
	"fmt"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type AddUserForm struct {
	Username string `form:"username" json:"username" valid:"Required;MaxSize(100)"`
	Password string `form:"password" json:"password" valid:"Required;MaxSize(100)"`
}

func UserRegister(c *gin.Context) {

	var addUserForm AddUserForm
	c.BindJSON(&addUserForm)

	valid := validation.Validation{}
	// a := auth{Username: a.Username, Password: a.Password}
	ok, _ := valid.Valid(&addUserForm)
	if !ok {
		code := e.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
		return
	}
	user := model.User{Username: addUserForm.Username, Password: addUserForm.Password}
	user.Save()

	code := e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

func UserInfo(c *gin.Context) {
	value, exists := c.Get(common.USER_NAME_Key)
	// value1, _ := c.Get(common.USER_ID_Key)

	username := ""

	if exists {
		username = value.(string)
	}

	userInfo := model.User{
		Username: username,
	}
	code := e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": userInfo,
	})
}

func DeleteUser(c *gin.Context) {
	fmt.Println("dfdfsd")
	id := c.Param("id")
	code := e.SUCCESS

	model.DeleteUser(id)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}
