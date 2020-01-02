package routers

import (
	"blog/controllers"
	"blog/pkg/setting"
	"blog/routers/api"
	v1 "blog/routers/api/v1"

	"blog/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.Use(middleware.Options)

	r.Use(
		middleware.JWTHeadAuth(middleware.AllowPathPrefixSkipper(
			middleware.JoinRouter("POST", "/api/v1/users"),
			middleware.JoinRouter("POST", "/api/v1/auth"),
		)))

	gin.SetMode(setting.RunMode)

	r.POST("/api/v1/auth/login", api.Login)

	apiv1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)

		// 用户列表
		// apiv1.GET("/users", controllers.AccountCtl)
		userCtl := controllers.UserController{}
		apiv1.GET("/users", userCtl.List)
		// 用户注册
		apiv1.POST("/users", userCtl.Create)
		// 获取用户信息
		apiv1.GET("/user-info", v1.UserInfo)
		// 删除用户
		apiv1.POST("/users/:id", v1.DeleteUser)

		//角色管理
		apiv1.GET("/roles", v1.GetRoles)
		// 新增角色
		apiv1.POST("/roles", v1.AddRole)

		//menu list
		apiv1.GET("/menus", v1.GetMenus)

		//add menu
		apiv1.POST("/menus", v1.AddMenu)
	}

	return r
}
