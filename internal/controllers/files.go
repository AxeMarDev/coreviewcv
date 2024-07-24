package controllers

import (
	"backendcv/internal/database"
	"backendcv/internal/models"
	"database/sql"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type File struct {
	FileName string `json:"file_name"`
	Mimitype string `json:"mime_type"`
	File     string `json:"file"`
}

func GetFiles(c *gin.Context) {

	var db *sql.DB
	db = database.Db

	projectID := c.Param("project_id")

	rows, err := db.Query("SELECT image_id, project_id, file_name, mime_type, size, created_at, updated_at FROM files WHERE project_id = ($1) ORDER BY image_id ASC ", projectID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query files"})
		return
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not complete request"})
		}
	}(rows)

	var files []models.File

	for rows.Next() {
		var p models.File
		if err := rows.Scan(&p.ImageId, &p.ProjectId, &p.Filename, &p.Mimetype, &p.Size, &p.CreatedAt, &p.UpdatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan files"})
			return
		}
		files = append(files, p)
	}
	fmt.Println(files)

	c.IndentedJSON(http.StatusOK, files)

}

func GetFile(c *gin.Context) {

	companyId, _ := c.Get("company_id")
	projectID := c.Param("project_id")
	fileID := c.Param("file_id")
	Mimetype := c.Query("mime_type")

	fmt.Println("step 1")

	fmt.Println("step 3")
	// Configure to use DigitalOcean Spaces
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("nyc3"), // Change to your region
		Credentials: credentials.NewStaticCredentials("DO00QFDGHD92FU6862XU", "frMi8x99jumgGfhW32qeCfw090dzWdqPlrmNb/mrRD4", ""),
		Endpoint:    aws.String("https://nyc3.digitaloceanspaces.com"),
	}))
	svc := s3.New(sess)

	fmt.Println("step 4")
	imageName := fmt.Sprintf("%g/%s/%s/%s", companyId, projectID, Mimetype, fileID)

	resp, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("coreview"), // Change to your space's name
		Key:    aws.String(imageName),  // Define the name of the file
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve file"})
		return
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	// Convert the bytes to a base64 string
	base64String := base64.StdEncoding.EncodeToString(bodyBytes)

	c.Stream(func(w io.Writer) bool {
		_, err := w.Write([]byte(fmt.Sprintf(`{"file":"%s"}`, base64String)))
		return err != nil // Return false to stop streaming
	})
}

func AddFile(c *gin.Context) {

	var db *sql.DB
	db = database.Db

	projectID := c.Param("project_id")
	companyId, _ := c.Get("company_id")

	var file File

	// Bind the received JSON to newPerson
	if err := c.ShouldBindJSON(&file); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(file.FileName)
	fmt.Println(file.Mimitype)

	// Decode base64 string to []byte
	data, err := base64.StdEncoding.DecodeString(strings.Split(file.File, "base64,")[1])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode base64 string"})
		return
	}

	// Configure to use DigitalOcean Spaces
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("nyc3"), // Change to your region
		Credentials: credentials.NewStaticCredentials("DO00QFDGHD92FU6862XU", "frMi8x99jumgGfhW32qeCfw090dzWdqPlrmNb/mrRD4", ""),
		Endpoint:    aws.String("https://nyc3.digitaloceanspaces.com"),
	}))

	query := `INSERT INTO files (project_id, file_name, mime_type, size) VALUES ($1, $2, $3, $4) RETURNING image_id`
	var id int
	err = db.QueryRow(query, projectID, file.FileName, file.Mimitype, len(data)).Scan(&id)

	if err != nil {
		log.Printf("Error while inserting new project: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add new project"})
		return
	}

	// Create an S3 service client
	svc := s3.New(sess)

	imageName := fmt.Sprintf("%g/%s/%s/%d", companyId, projectID, file.Mimitype, id)

	// Upload the file to the specific space
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String("coreview"), // Change to your space's name
		Key:    aws.String(imageName),  // Define the name of the file
		Body:   strings.NewReader(string(data)),
		ACL:    aws.String("public-read"), // or another ACL policy
	})

	if err != nil {
		log.Println("Failed to upload", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file to space"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})

}
