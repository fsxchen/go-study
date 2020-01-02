package main

import (
	"blog/middleware"
	"blog/model"
	"blog/pkg/setting"
	"blog/routers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "blog/docs"
)

type Login struct {
	ID       uint   `gorm:"primary_key"`
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host http://localhost:8081

func main() {
	router := gin.Default()
	model.InitDB()
	routersInit := routers.InitRouter()
	// account.Name = "hello"
	// account.Email = "hello@hello.com"
	// account.Save()

	router.Use(middleware.TestMiddle())
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        routersInit,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
	// app.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "NAME_OF_ENV_VARIABLE"))

	// app.GET("/string/:name", controllers.AccountCtl)
	// app.POST("/string", controllers.RegisterCtl)

	// app.POST("/role", controllers.AddRole)

	// app.Run(":8081")
}
