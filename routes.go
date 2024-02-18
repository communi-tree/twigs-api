package main

import (
	v1 "github.com/communi-tree/twigs-api/app/controllers"
	"github.com/communi-tree/twigs-api/app/utils/middleware"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/login", v1.LoginHandler)
	router.POST("/create_user", v1.CreateUser)

	router.Use(middleware.AuthMiddelware())

	// router.GET("/users", v1.UserIndex)
	router.GET("/user/:id", v1.UserShow)
	router.POST("/subdivision", v1.CreateSubdivision)
	return router
}
