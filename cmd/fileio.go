package main

import (
	"encoding/json"
	"os"
)

const PostsDB = "db/posts.json"

func SavePostToDB(post []Post) error {
	data, err := json.Marshal(post)
	if err != nil {
		return err
	}
	return os.WriteFile(PostsDB, data, 0644)
}

func LoadPostFromDB() ([]Post, error) {
	data, err := os.ReadFile(PostsDB)
	if err != nil {
		return nil, err
	}
	var posts []Post
	err = json.Unmarshal(data, &posts)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
