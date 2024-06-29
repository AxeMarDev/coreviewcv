package controllers

import (
	"backendcv/internal/database"
	"backendcv/internal/models"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func GetClients(c *gin.Context) {

	var db *sql.DB
	db = database.Db

	companyId, _ := c.Get("company_id")

	fmt.Println(companyId)

	rows, err := db.Query("SELECT id, name, username, email, phone, company_id FROM client WHERE company_id = ($1) ORDER BY id ASC ", companyId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query clients"})
		return
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not complete request"})
		}
	}(rows)

	var clients []models.Client

	for rows.Next() {
		var p models.Client
		if err := rows.Scan(&p.ID, &p.Name, &p.Username, &p.Email, &p.Phone, &p.Companyid); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan person"})
			return
		}
		clients = append(clients, p)
	}
	fmt.Println(clients)

	c.IndentedJSON(http.StatusOK, clients)

}

func AddClient(c *gin.Context) {
	var newClient models.Client

	var db *sql.DB
	db = database.Db

	companyId, _ := c.Get("company_id")

	fmt.Println(companyId)

	// Bind the received JSON to newPerson
	if err := c.ShouldBindJSON(&newClient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newClient.Hashpassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not hash passwword"})
	}

	// Insert newPerson into the database
	query := `INSERT INTO client (name, username, email, hash_password, phone, company_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	var id int
	err = db.QueryRow(query, newClient.Name, newClient.Username, newClient.Email, hashedPassword, newClient.Phone, companyId).Scan(&id)

	if err != nil {
		log.Printf("Error while inserting new person: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add new person"})
		return
	}

	// Return the new person as JSON
	c.JSON(http.StatusCreated, newClient)
}

func DeleteClient(c *gin.Context) {

	var db *sql.DB
	db = database.Db

	id := c.Query("id")

	// Execute the delete query
	query := `DELETE FROM client WHERE id = $1`
	result, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Error while deleting client: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete client"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking deletion result"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No client found with the provided ID"})
		return
	}

	// Return success message
	c.JSON(http.StatusOK, gin.H{"message": "client deleted successfully"})

}
