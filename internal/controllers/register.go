package controllers

import (
	"backendcv/internal/database"
	"database/sql"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"

	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type Mastercompanyemployee struct {
	ID          string `json:"id"`
	Companyname string `json:"company_name"`
	Companycode string `json:"company_code"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	Phone       string `json:"phone"`
	Password    string `json:"password"`
}

func Register(c *gin.Context) {
	var newCompany Mastercompanyemployee

	var db *sql.DB
	db = database.Db

	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "fatal"})
		return
	}

	fmt.Println("step 1")

	// Bind the received JSON to newPerson
	if err := c.ShouldBindJSON(&newCompany); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("step 2")

	// Insert newPerson into the database
	query := `INSERT INTO company (company_name, company_code) VALUES ($1, $2) RETURNING id`
	var company_id int
	err := db.QueryRow(query, newCompany.Companyname, newCompany.Companycode).Scan(&company_id)

	fmt.Println(company_id)

	if err != nil {
		log.Printf("Error while inserting new person: %v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add new company"})
		return
	}

	fmt.Println("sign up password")
	fmt.Println(newCompany.Password)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newCompany.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not hash passwword"})
	}

	var idNew int
	queryMasterEmployee := `INSERT INTO employee (name, username, email, hash_password, phone, isadmin, company_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	errEmployee := db.QueryRow(queryMasterEmployee, "ADMIN", newCompany.Username, newCompany.Email, string(hashedPassword), newCompany.Phone, true, company_id).Scan(&idNew)

	if errEmployee != nil {
		log.Printf("Error while inserting new person: %v", errEmployee.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add master Employee"})
		return
	}

	_, err = db.Exec(`UPDATE company SET masteremployee_id = $1 WHERE id = $2`, idNew, company_id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not return auth token"})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = newCompany.Username
	claims["company_id"] = company_id
	claims["exp"] = time.Now().Add(time.Hour * 1000).Unix()

	t, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not return auth token"})
	}

	// Return the new person as JSON
	c.JSON(http.StatusCreated, gin.H{"id": newCompany.Username, "company_name": company_id, "jwt": t})
}
