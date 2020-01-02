package v1

import (
	"blog/controllers/common"
	"blog/model"
	"log"
	"net/http"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"blog/pkg/e"
	"blog/pkg/setting"
	"blog/pkg/util"
)

//获取单个文章
func GetArticle(c *gin.Context) {
	id := c.Param("id")

	// valid := validation.Validation{}
	// valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	// var data interface{}
	// if !valid.HasErrors() {

	// } else {
	// 	for _, err := range valid.Errors {
	// 		log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
	// 	}
	// }

	var data model.Article
	if model.ExistArticleByID(id) {
		data = model.GetArticle(id)
		code = e.SUCCESS
	} else {
		code = e.ERROR_NOT_EXIST_ARTICLE
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//获取多个文章
func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state

		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	var tagId int = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagId

		valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	}

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS

		data["lists"] = model.GetArticles(util.GetPage(c), setting.PageSize, maps)
		data["total"] = model.GetArticleTotal(maps)

	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//新增文章

type AddArticleForm struct {
	TagID           int    `form:"tag_id" json:"tag_id" valid:"Required;Min(1)"`
	Title           string `form:"title" json:"title" valid:"Required;MaxSize(100)"`
	Desc            string `form:"desc" json:"desc" valid:"Required;MaxSize(255)"`
	Content         string `form:"content" json:"content" valid:"Required;MaxSize(65535)"`
	ContentMarkdown string `form:"content_markdown" json:"content_markdown" valid:"Required;MaxSize(65535)"`
	CoverImageUrl   string `form:"cover_image_url" json:"cover_image_url" valid:"Required;MaxSize(255)"`
	State           int    `form:"state" json:"state" valid:"Range(0,1)"`
	CreatedBy       string
}

func AddArticle(c *gin.Context) {
	// TODO check args
	// valid := validation.Validation{}
	// valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	// valid.Required(title, "title").Message("标题不能为空")
	// valid.Required(desc, "desc").Message("简述不能为空")
	// valid.Required(content, "content").Message("内容不能为空")
	// valid.Required(createdBy, "created_by").Message("创建人不能为空")
	// valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	var addArticleForm AddArticleForm
	// code := e.INVALID_PARAMS
	// if !valid.HasErrors() {
	// 	if res, _ := models.ExistTagByID(tagId); !res {
	// 		data := make(map[string]interface{})
	// 		data["tag_id"] = tagId
	// 		data["title"] = title
	// 		data["desc"] = desc
	// 		data["content"] = content
	// 		data["created_by"] = createdBy
	// 		data["state"] = state

	// 		models.AddArticle(data)
	// 		code = e.SUCCESS
	// 	} else {
	// 		code = e.ERROR_NOT_EXIST_TAG
	// 	}
	// } else {
	// 	for _, err := range valid.Errors {
	// 		log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
	// 	}
	// }

	c.BindJSON(&addArticleForm)

	value, isExists := c.Get(common.USER_ID_Key)
	if isExists {
		addArticleForm.CreatedBy = value.(string)
	}

	article := model.Article{
		Title:           addArticleForm.Title,
		CreatedBy:       addArticleForm.CreatedBy,
		Content:         addArticleForm.Content,
		ContentMarkdown: addArticleForm.ContentMarkdown,
	}

	article.Save()

	code := e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

//修改文章
func EditArticle(c *gin.Context) {
	valid := validation.Validation{}

	id := c.Param("id")
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if model.ExistArticleByID(id) {
			if res, _ := model.ExistTagByID(tagId); res {
				data := make(map[string]interface{})
				if tagId > 0 {
					data["tag_id"] = tagId
				}
				if title != "" {
					data["title"] = title
				}
				if desc != "" {
					data["desc"] = desc
				}
				if content != "" {
					data["content"] = content
				}

				data["modified_by"] = modifiedBy

				model.EditArticle(id, data)
				code = e.SUCCESS
			} else {
				code = e.ERROR_NOT_EXIST_TAG
			}
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

//删除文章
func DeleteArticle(c *gin.Context) {
	id := c.Param("id")

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if model.ExistArticleByID(id) {
			model.DeleteArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
