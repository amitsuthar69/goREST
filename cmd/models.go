package main

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
