// controllers/post.go
package controllers

import (
	"net/http"

	"github.com/billbull21/go-crud.git/models"
	"github.com/gin-gonic/gin"
)

type CreatePostInput struct {
    Title   	string `json:"title" binding:"required"`
    Description	string `json:"description" binding:"required"`
    Published	bool `json:"published" binding:"required"`
}

// CREATE
func CreatePost(c *gin.Context) {
    var input CreatePostInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    post := models.Post{
		Title: input.Title, 
		Description: input.Description,
		Published: input.Published,
	}
    models.DB.Create(&post)

    c.JSON(http.StatusOK, gin.H{"data": post})
}

// UPDATED
type UpdatePostInput struct {
    Title   	string `json:"title" binding:"required"`
    Description	string `json:"description" binding:"required"`
    Published	bool `json:"published" binding:"required"`
}

func UpdatePost(c *gin.Context) {
    var post models.Post
    if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
        return
    }

    var input UpdatePostInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    updatedPost := models.Post{
        Title: input.Title, 
        Description: input.Description,
        Published: input.Published,
    }

    models.DB.Model(&post).Updates(&updatedPost)
    c.JSON(http.StatusOK, gin.H{"data": post})
}

// GET ALL
func FindPosts(c *gin.Context) {
    var posts []models.Post
    models.DB.Find(&posts)

    c.JSON(http.StatusOK, gin.H{"data": posts})
}

// GET ALL BY ID
func FindPost(c *gin.Context) {
    var post models.Post

    if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": post})
}

// Delete post by ID
func DeletePost(c *gin.Context) {
    var post models.Post
    if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
        return
    }

    models.DB.Delete(&post)
    c.JSON(http.StatusOK, gin.H{"data": "success"})
}