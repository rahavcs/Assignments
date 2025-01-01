package models

import "time"

// Blog struct represents a blog post.
type Blog struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	Timestamp time.Time `json:"timestamp"`
}
