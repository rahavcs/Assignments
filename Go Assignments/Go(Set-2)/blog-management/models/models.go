package models

import (
	"blog-management/database"
)

type Blog struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Author    string `json:"author"`
	Timestamp string `json:"timestamp"`
}

func GetAllBlogs() ([]Blog, error) {
	rows, err := database.DB.Query("SELECT id, title, content, author, timestamp FROM blogs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []Blog
	for rows.Next() {
		var blog Blog
		if err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.Timestamp); err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}
	return blogs, nil
}

func CreateBlog(blog Blog) error {
	_, err := database.DB.Exec("INSERT INTO blogs (title, content, author) VALUES (?, ?, ?)", blog.Title, blog.Content, blog.Author)
	return err
}

func GetBlogByID(id int) (Blog, error) {
	var blog Blog
	err := database.DB.QueryRow("SELECT id, title, content, author, timestamp FROM blogs WHERE id = ?", id).Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.Timestamp)
	return blog, err
}

func UpdateBlog(id int, blog Blog) error {
	_, err := database.DB.Exec("UPDATE blogs SET title = ?, content = ?, author = ? WHERE id = ?", blog.Title, blog.Content, blog.Author, id)
	return err
}

func DeleteBlog(id int) error {
	_, err := database.DB.Exec("DELETE FROM blogs WHERE id = ?", id)
	return err
}
