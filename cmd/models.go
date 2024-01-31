package main

type Post struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	AuthorName string `json:"authorName"`
}

// let's just cope with Posts as DB items.
var Posts = []Post{}
