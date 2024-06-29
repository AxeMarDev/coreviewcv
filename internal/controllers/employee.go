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

func GetEmployees(c *gin.Context) {

	var db *sql.DB
	db = database.Db

	companyId, _ := c.Get("company_id")

	fmt.Println(companyId)

	rows, err := db.Query("SELECT id, name, username, email, phone, company_id FROM employee WHERE company_id = ($1) ORDER BY id ASC ", companyId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query employees"})
		return
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not complete request"})
		}
	}(rows)

	var employees []models.Employee

	for rows.Next() {
		var p models.Employee
		if err := rows.Scan(&p.ID, &p.Name, &p.Username, &p.Email, &p.Phone, &p.Companyid); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan person"})
			return
		}
		employees = append(employees, p)
	}
	fmt.Println(employees)

	c.IndentedJSON(http.StatusOK, employees)
}

func AddEmployees(c *gin.Context) {

	var newEmployee models.Employee

	var db *sql.DB
	db = database.Db

	companyId, _ := c.Get("company_id")

	fmt.Println(companyId)

	// Bind the received JSON to newPerson
	if err := c.ShouldBindJSON(&newEmployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newEmployee.Hashpassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not hash passwword"})
	}

	// Insert newPerson into the database
	query := `INSERT INTO employee (name, username, email, hash_password, phone, company_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	var id int
	err = db.QueryRow(query, newEmployee.Name, newEmployee.Username, newEmployee.Email, hashedPassword, newEmployee.Phone, companyId).Scan(&id)

	if err != nil {
		log.Printf("Error while inserting new employee: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add new employee"})
		return
	}

	// Return the new person as JSON
	c.JSON(http.StatusCreated, newEmployee)
}
