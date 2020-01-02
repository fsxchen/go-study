package v1

import (
	"fmt"
	"net/http"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"blog/model"
	"blog/pkg/e"
	"blog/pkg/setting"
	"blog/pkg/util"
)

//获取多个文章标签
// @Summary Get multiple article tags
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags [get]
func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	fmt.Println(maps)

	code := e.SUCCESS
	var err error
	data["lists"], err = model.GetTags(util.GetPage(c), setting.PageSize, maps)
	fmt.Println(err)
	fmt.Println("----------------")
	data["total"], _ = model.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

type AddTagForm struct {
	Name      string `form:"name" json:"name" valid:"Required;MaxSize(100)"`
	CreatedBy string `form:"created_by" json:"created_by" valid:"Required;MaxSize(100)"`
	State     int    `form:"state" json:"state" valid:"Range(0,1)"`
}

// @Summary Add article tag
// @Produce  json
// @Param name body string true "Name"
// @Param state body int false "State"
// @Param created_by body int false "CreatedBy"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags [post]
func AddTag(c *gin.Context) {
	var addTagForm AddTagForm
	c.BindJSON(&addTagForm)

	// TODO validation

	// valid := validation.Validation{}
	// valid.Required(name, "name").Message("名称不能为空")
	// valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	// valid.Required(createdBy, "created_by").Message("创建人不能为空")
	// valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	// valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.SUCCESS

	model.AddTag(addTagForm.Name, addTagForm.State, addTagForm.CreatedBy)

	// for _, i := range valid.Errors {
	// 	fmt.Println(i.Message)
	// }
	// if !valid.HasErrors() {
	// 	if res, _ := model.ExistTagByName(name); !res {
	// 		code = e.SUCCESS
	// 		model.AddTag(name, state, createdBy)
	// 	} else {
	// 		code = e.ERROR_EXIST_TAG
	// 	}
	// }

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

//修改文章标签
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if res, _ := model.ExistTagByID(id); res {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}

			model.EditTag(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

//删除文章标签
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if res, _ := model.ExistTagByID(id); res {
			model.DeleteTag(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// type EditTagForm struct {
// 	ID         int    `form:"id" valid:"Required;Min(1)"`
// 	Name       string `form:"name" valid:"Required;MaxSize(100)"`
// 	ModifiedBy string `form:"modified_by" valid:"Required;MaxSize(100)"`
// 	State      int    `form:"state" valid:"Range(0,1)"`
// }

// // @Summary Update article tag
// // @Produce  json
// // @Param id path int true "ID"
// // @Param name body string true "Name"
// // @Param state body int false "State"
// // @Param modified_by body string true "ModifiedBy"
// // @Success 200 {object} app.Response
// // @Failure 500 {object} app.Response
// // @Router /api/v1/tags/{id} [put]
// func EditTag(c *gin.Context) {
// 	var (
// 		appG = app.Gin{C: c}
// 		form = EditTagForm{ID: com.StrTo(c.Param("id")).MustInt()}
// 	)

// 	httpCode, errCode := app.BindAndValid(c, &form)
// 	if errCode != e.SUCCESS {
// 		appG.Response(httpCode, errCode, nil)
// 		return
// 	}

// 	tagService := tag_service.Tag{
// 		ID:         form.ID,
// 		Name:       form.Name,
// 		ModifiedBy: form.ModifiedBy,
// 		State:      form.State,
// 	}

// 	exists, err := tagService.ExistByID()
// 	if err != nil {
// 		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
// 		return
// 	}

// 	if !exists {
// 		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
// 		return
// 	}

// 	err = tagService.Edit()
// 	if err != nil {
// 		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_TAG_FAIL, nil)
// 		return
// 	}

// 	appG.Response(http.StatusOK, e.SUCCESS, nil)
// }

// // @Summary Delete article tag
// // @Produce  json
// // @Param id path int true "ID"
// // @Success 200 {object} app.Response
// // @Failure 500 {object} app.Response
// // @Router /api/v1/tags/{id} [delete]
// func DeleteTag(c *gin.Context) {
// 	appG := app.Gin{C: c}
// 	valid := validation.Validation{}
// 	id := com.StrTo(c.Param("id")).MustInt()
// 	valid.Min(id, 1, "id").Message("ID必须大于0")

// 	if valid.HasErrors() {
// 		app.MarkErrors(valid.Errors)
// 		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
// 	}

// 	tagService := tag_service.Tag{ID: id}
// 	exists, err := tagService.ExistByID()
// 	if err != nil {
// 		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
// 		return
// 	}

// 	if !exists {
// 		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
// 		return
// 	}

// 	if err := tagService.Delete(); err != nil {
// 		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_TAG_FAIL, nil)
// 		return
// 	}

// 	appG.Response(http.StatusOK, e.SUCCESS, nil)
// }

// // @Summary Export article tag
// // @Produce  json
// // @Param name body string false "Name"
// // @Param state body int false "State"
// // @Success 200 {object} app.Response
// // @Failure 500 {object} app.Response
// // @Router /api/v1/tags/export [post]
// func ExportTag(c *gin.Context) {
// 	appG := app.Gin{C: c}
// 	name := c.PostForm("name")
// 	state := -1
// 	if arg := c.PostForm("state"); arg != "" {
// 		state = com.StrTo(arg).MustInt()
// 	}

// 	tagService := tag_service.Tag{
// 		Name:  name,
// 		State: state,
// 	}

// 	filename, err := tagService.Export()
// 	if err != nil {
// 		appG.Response(http.StatusInternalServerError, e.ERROR_EXPORT_TAG_FAIL, nil)
// 		return
// 	}

// 	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
// 		"export_url":      export.GetExcelFullUrl(filename),
// 		"export_save_url": export.GetExcelPath() + filename,
// 	})
// }

// // @Summary Import article tag
// // @Produce  json
// // @Param file body file true "Excel File"
// // @Success 200 {object} app.Response
// // @Failure 500 {object} app.Response
// // @Router /api/v1/tags/import [post]
// func ImportTag(c *gin.Context) {
// 	appG := app.Gin{C: c}

// 	file, _, err := c.Request.FormFile("file")
// 	if err != nil {
// 		logging.Warn(err)
// 		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
// 		return
// 	}

// 	tagService := tag_service.Tag{}
// 	err = tagService.Import(file)
// 	if err != nil {
// 		logging.Warn(err)
// 		appG.Response(http.StatusInternalServerError, e.ERROR_IMPORT_TAG_FAIL, nil)
// 		return
// 	}

// 	appG.Response(http.StatusOK, e.SUCCESS, nil)
// }
