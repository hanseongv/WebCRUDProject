package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/users", GetUsers)
	router.GET("/users/:id", GetUser)
	router.POST("/users", CreateUser)
	router.PUT("/users/:id", UpdateUser)
	router.DELETE("/users/:id", DeleteUser)
	router.Run(":8080")
}
