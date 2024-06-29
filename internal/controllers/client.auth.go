package controllers

import (
	"backendcv/internal/database"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func ClientAuth(c *gin.Context) {
	var newAuth Authenticate

	var db *sql.DB
	db = database.Db

	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "fatal"})
		return
	}

	fmt.Println("step 2")
	// Bind the received JSON to newPerson
	if err := c.ShouldBindJSON(&newAuth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("login password")
	fmt.Println(newAuth.Password)

	fmt.Println("step 3")
	query := `SELECT hash_password, company_id, id FROM client WHERE username = $1`
	var dbPassword string
	var companyID int
	var userID int

	// Execute the query
	err := db.QueryRow(query, newAuth.Username).Scan(&dbPassword, &companyID, &userID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(newAuth.Password)) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "incorrect password or email"})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = newAuth.Username
	claims["company_id"] = companyID
	claims["account_type"] = "client"
	claims["exp"] = time.Now().Add(time.Hour * 1000).Unix()

	t, err := token.SignedString([]byte("your_secret_key"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not return auth token"})
		return
	}

	// Return the new person as JSON
	c.JSON(http.StatusCreated, gin.H{"id": newAuth.Username, "company_name": companyID, "jwt": t})
}
