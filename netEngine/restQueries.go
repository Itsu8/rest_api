package net_engine

import (
	"github.com/Itsu8/rest_api/controllers"
	"github.com/gin-gonic/gin"
)

func RunNetEngine() {
	router := gin.Default()

	router.POST("/posts", controllers.PostsCreate)
	router.GET("/posts", controllers.PostsRead)
	router.GET("/posts/:id", controllers.PostsShow)
	router.PUT("/posts/:id", controllers.PostsUpdate)

	router.Run()
}
