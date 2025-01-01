package handlers

import (
	"blog-management/config"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBlog(c *gin.Context) {
	var blog struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		Author  string `json:"author"`
	}

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := config.DB.Exec("INSERT INTO blogs (title, content, author) VALUES (?, ?, ?)", blog.Title, blog.Content, blog.Author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	lastInsertID, _ := result.LastInsertId()
	c.JSON(http.StatusCreated, gin.H{"id": lastInsertID})
}

func GetBlogByID(c *gin.Context) {
	id := c.Param("id")
	var blog struct {
		ID        int    `json:"id"`
		Title     string `json:"title"`
		Content   string `json:"content"`
		Author    string `json:"author"`
		Timestamp string `json:"timestamp"`
	}

	row := config.DB.QueryRow("SELECT * FROM blogs WHERE id = ?", id)
	if err := row.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.Timestamp); err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, blog)
}

func GetAllBlogs(c *gin.Context) {
	rows, err := config.DB.Query("SELECT * FROM blogs")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var blogs []struct {
		ID        int    `json:"id"`
		Title     string `json:"title"`
		Content   string `json:"content"`
		Author    string `json:"author"`
		Timestamp string `json:"timestamp"`
	}

	for rows.Next() {
		var blog struct {
			ID        int    `json:"id"`
			Title     string `json:"title"`
			Content   string `json:"content"`
			Author    string `json:"author"`
			Timestamp string `json:"timestamp"`
		}
		if err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.Timestamp); err != nil {
			log.Fatal("Failed to scan row: ", err)
		}
		blogs = append(blogs, blog)
	}

	c.JSON(http.StatusOK, blogs)
}

func UpdateBlog(c *gin.Context) {
	id := c.Param("id")
	var blog struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		Author  string `json:"author"`
	}

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := config.DB.Exec("UPDATE blogs SET title = ?, content = ?, author = ? WHERE id = ?", blog.Title, blog.Content, blog.Author, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog updated"})
}

func DeleteBlog(c *gin.Context) {
	id := c.Param("id")
	_, err := config.DB.Exec("DELETE FROM blogs WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog deleted"})
}
