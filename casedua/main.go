package main

import (
	"encoding/json"
	"os"
)

type Comment struct {
	CommentID      int       `json:"commentId"`
	CommentContent string    `json:"commentContent"`
	Replies        []Comment `json:"replies,omitempty"`
}

func main() {
	b, err := os.ReadFile("comments.json")
	if err != nil {
		panic(err)
	}

	var comments []Comment
	if err := json.Unmarshal(b, &comments); err != nil {
		panic(err)
	}

	total := count(comments)
	println("Total comments: ", total)
}

func count(c []Comment) int {
	total := len(c)
	for _, comment := range c {
		total += count(comment.Replies)
	}
	return total
}
