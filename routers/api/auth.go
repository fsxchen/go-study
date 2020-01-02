package api

import (
	"blog/model"
	"blog/pkg/setting"
	"fmt"
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"blog/pkg/e"
	"blog/pkg/util"
)

type AuthForm struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func Login(c *gin.Context) {

	var authForm AuthForm

	c.BindJSON(&authForm)

	valid := validation.Validation{}
	// a := auth{Username: a.Username, Password: a.Password}
	ok, _ := valid.Valid(&authForm)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		user, isExist := model.UserLogin(authForm.Username, authForm.Password)
		fmt.Println(isExist)
		if isExist {
			token, err := util.GenerateToken(user.Username, user.Password, user.ID.String(), setting.SessionTTL)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}

		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
