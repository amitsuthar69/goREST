package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CREATE
func AddPost(c *gin.Context) {
	var newPost Post
	if err := c.BindJSON(&newPost); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Posts = append(Posts, newPost)

	if err := SavePostToDB(Posts); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to save into db"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newPost)
}

// READ
func GetPosts(c *gin.Context) {
	posts, err := LoadPostFromDB()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to load from db"})
		return
	}
	c.IndentedJSON(http.StatusOK, posts)
}

func GetPostById(c *gin.Context) {
	id := c.Param("id")
	for _, post := range Posts {
		if post.Id == id {
			c.IndentedJSON(http.StatusOK, post)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found!"})
}

// UPDATE
func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var updatedPost Post
	if err := c.BindJSON(&updatedPost); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, post := range Posts {
		if post.Id == id {
			Posts[i] = updatedPost
			if err := SavePostToDB(Posts); err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to save into db"})
				return
			}
			c.IndentedJSON(http.StatusOK, Posts[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found!"})
}

// DELETE
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	for i, post := range Posts {
		if post.Id == id {
			Posts = append(Posts[:i], Posts[i+1:]...)
			if err := SavePostToDB(Posts); err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to save into db"})
				return
			}
			c.IndentedJSON(http.StatusOK, gin.H{"message": "post deleted successfully"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found!"})
}
