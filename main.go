package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Post struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	AuthorName string `json:"authorName"`
}

// let's just cope with Posts as DB items.
var Posts = []Post{
	{Id: "1", Title: "Post1", AuthorName: "amit"},
	{Id: "2", Title: "Post2", AuthorName: "rishabh"},
	{Id: "3", Title: "Post3", AuthorName: "sumit"},
	{Id: "4", Title: "Post4", AuthorName: "harshil"},
}

// CREATE
func addPost(c *gin.Context) {
	var newPost Post
	if err := c.BindJSON(&newPost); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Posts = append(Posts, newPost)
	c.IndentedJSON(http.StatusCreated, newPost)
}

// READ
func getPosts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Posts)
}

func getPostById(c *gin.Context) {
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
func updatePost(c *gin.Context) {
	id := c.Param("id")
	var updatedPost Post
	if err := c.BindJSON(&updatedPost); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, post := range Posts {
		if post.Id == id {
			Posts[i] = updatedPost
			c.IndentedJSON(http.StatusOK, Posts[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found!"})
}

// DELETE
func deletePost(c *gin.Context) {
	id := c.Param("id")
	for i, post := range Posts {
		if post.Id == id {
			Posts = append(Posts[:i], Posts[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "post deleted successfully"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found!"})
}

func main() {
	router := gin.Default()
	router.GET("/posts", getPosts)
	router.POST("/posts", addPost)
	router.GET("/posts/:id", getPostById)
	router.PUT("/posts/:id", updatePost)
	router.DELETE("/posts/:id", deletePost)
	router.Run("localhost:3000")
}

/*
curl http://localhost:3000/posts \
--include \
--header "Contnt-Type: application/json" \
--request "POST" \
--data '{"id":3, "title":"Post3","authorName":"amit"}'

curl http://localhost:3000/posts/1 \
--include \
--header "Content-Type: application/json" \
--request "PUT" \
--data '{"id":"1", "title":"UpdatedPost1", "authorName":"sumit"}'

curl -X DELETE http://localhost:3000/posts/2
*/
