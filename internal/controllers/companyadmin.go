package controllers

import (
	"backendcv/internal/database"
	"backendcv/internal/models"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

func AdminRegister(c *gin.Context) {

	var db *sql.DB
	db = database.Db

	var newAdmin models.CompanyAdmin

	// Bind the received JSON to newPerson
	if err := c.ShouldBindJSON(&newAdmin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newAdmin.Hashpassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not hash password"})
		return // added return here to stop execution on error
	}

	var id int
	queryMasterEmployee := `INSERT INTO companyadmin (name, username, email, hash_password) VALUES ($1, $2, $3, $4) RETURNING id`
	errEmployee := db.QueryRow(queryMasterEmployee, newAdmin.Name, newAdmin.Username, newAdmin.Email, string(hashedPassword)).Scan(&id)
	if errEmployee != nil {
		log.Printf("Error while inserting new admin: %v", errEmployee)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add master Employee"})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = newAdmin.Username
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 1000).Unix()

	t, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not return auth token"})
	}

	// Return the new person as JSON
	c.JSON(http.StatusCreated, gin.H{"id": newAdmin.Username, "type": "admin", "jwt": t})
}
