package main

import (
	"github.com/gin-gonic/gin"
	"github.com/repository/controllers"
	"github.com/repository/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/", controllers.PostGetAll)
	r.GET("/post/:id", controllers.PostGet)
	r.POST("/post", controllers.PostCreate)
	r.PATCH("/post", controllers.PostUpdate)
	r.DELETE("/post/:id", controllers.PostDelete)

	r.Run()
}
