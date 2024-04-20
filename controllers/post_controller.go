package controllers

import (
	data "go-side-project/initializers"
	model "go-side-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Post CRUD
func GetAllPosts(c *gin.Context){
	var posts []model.Post
    if err := data.DB.Find(&posts).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func CreatePost(c *gin.Context){
	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request data"})
	}

	if err := data.DB.Create(&post).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Create post successfully"})
}

func ShowPost(c *gin.Context){
	id := c.Param("id")
	var post model.Post

	if err := data.DB.Find(&post, id).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{ "post": post,
	})
}

func UpdatePost(c *gin.Context) {
    id := c.Param("id")
    var post model.Post

    if err := data.DB.First(&post, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
        return
    }

    if err := c.ShouldBindJSON(&post); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request data"})
        return
    }

    if err := data.DB.Save(&post).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
}


func DeletePost(c *gin.Context){
	id := c.Param("id")
	if err := data.DB.Delete(&model.Post{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}
	c.JSON(http.StatusOK, gin.H{ "message": "Delete post successfully",
	})
}