package main

import (
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

func main() {
	router := gin.Default()
	router.GET("/posts", GetPosts)
	router.POST("/posts", AddPost)
	router.GET("/posts/:id", GetPostById)
	router.PUT("/posts/:id", UpdatePost)
	router.DELETE("/posts/:id", DeletePost)
	router.Run("localhost:3000")
}

/*
curl http://localhost:3000/posts \
--include \
--request "POST" \
--data '{"id":3, "title":"Post3","authorName":"amit"}'

curl http://localhost:3000/posts/1 \
--include \
--header "Content-Type: application/json" \
--request "PUT" \
--data '{"id":"1", "title":"UpdatedPost1", "authorName":"sumit"}'

curl -X DELETE http://localhost:3000/posts/2
*/
