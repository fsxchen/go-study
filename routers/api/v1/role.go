package v1

import (
	"blog/model"
	"blog/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRoles(c *gin.Context) {
	data, _ := model.GetRoles()
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

type AddRoleForm struct {
	Name     *string `form:"name" json:"name" valid:"Required;MaxSize(100)"`
	RecordID string  `json:"record_id"`
	Sequence *int    `json:"sequence"`
	Memo     *string `json:"memo"`
}

func AddRole(c *gin.Context) {

	var addRole AddRoleForm
	c.BindJSON(&addRole)
	code := e.SUCCESS

	role := model.Role{
		Name:     addRole.Name,
		RecordID: addRole.RecordID,
		Sequence: addRole.Sequence,
		Memo:     addRole.Memo,
	}

	role.Save()
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": role,
	})
}
