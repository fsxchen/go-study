package v1

import (
	"blog/model"
	"blog/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddMenuForm struct {
	MenuLevel  uint         `json:"-"`
	ParentId   string       `json:"parentId"`
	ParentPath string       `json:"parentPath"`
	Path       string       `json:"path"`
	Name       string       `json:"name"`
	Hidden     bool         `json:"hidden"`
	Component  string       `json:"component"`
	Router     string       `json:"router"`
	Icon       string       `json:"icon"`
	Meta       string       `json:"meta"`
	NickName   string       `json:"nickName"`
	Children   []model.Menu `json:"children"`
}

func AddMenu(c *gin.Context) {
	var addMenuForm AddMenuForm
	c.BindJSON(&addMenuForm)
	newMenu := model.Menu{
		Name: addMenuForm.Name,
		Path: addMenuForm.Path,
	}
	newMenu.Save()

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": newMenu,
	})
}

func GetMenus(c *gin.Context) {
	menus, _ := model.GetMenus()

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": menus,
	})
}
