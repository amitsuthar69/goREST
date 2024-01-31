package main

import (
	"github.com/gin-gonic/gin"
)

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
