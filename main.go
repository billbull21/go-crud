// main.go
package main

import (
	"github.com/billbull21/go-crud.git/controllers"
	"github.com/billbull21/go-crud.git/models"
	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    models.ConnectDatabase()

    // router
    router.POST("/posts", controllers.CreatePost)  // Created
    router.GET("/posts", controllers.FindPosts) // Get all
    router.GET("/posts/:id", controllers.FindPost)  // Get By Id
    router.PATCH("/posts/:id", controllers.UpdatePost) // Updated
    router.DELETE("/posts/:id", controllers.DeletePost) // Delete

    router.Run("localhost:8080")
}