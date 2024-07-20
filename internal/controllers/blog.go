package controllers

import (
	"backendcv/internal/database"
	"backendcv/internal/models"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func GetBlogs(c *gin.Context) {

	var db *sql.DB
	db = database.Db

	rows, err := db.Query("SELECT id, title, subtitle FROM blogs ORDER BY id ASC ")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query blogs"})
		return
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not complete request"})
		}
	}(rows)

	var blogs []models.Blog

	for rows.Next() {
		var p models.Blog
		if err := rows.Scan(&p.ID, &p.Title, &p.Subtitle); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan person"})
			return
		}
		blogs = append(blogs, p)
	}

	c.IndentedJSON(http.StatusOK, blogs)

}

func AddBlog(c *gin.Context) {
	var newBlog models.Blog

	Id, _ := c.Get("id")

	var db *sql.DB
	db = database.Db

	// Bind the received JSON to newPerson
	if err := c.ShouldBindJSON(&newBlog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert newPerson into the database
	query := `INSERT INTO blogs (title, subtitle, author_id, date_posted) VALUES ($1, $2, $3, $4) RETURNING id`
	var id int
	err := db.QueryRow(query, newBlog.Title, newBlog.Subtitle, Id, time.Now()).Scan(&id)

	if err != nil {
		log.Printf("Error while inserting new blog: %v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add new blog"})
		return
	}

	// Return the new person as JSON
	c.JSON(http.StatusCreated, newBlog)
}

func DeleteBlog(c *gin.Context) {

	var db *sql.DB //
	db = database.Db

	id := c.Param("blog_id")

	fmt.Println(id)

	// Execute the delete query
	query := `DELETE FROM blogs WHERE id = $1`
	result, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Error while deleting blog: %v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete blog"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking deletion result"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No blog found with the provided ID"})
		return
	}

	// Return success message
	c.JSON(http.StatusOK, gin.H{"message": "project deleted successfully"})

}
