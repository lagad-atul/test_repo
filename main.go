package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	db *gorm.DB
)

func init() {
	// Initialize the database connection
	var err error
	// Replace "atul:Choice123@tcp(localhost:3306)/leadapi" with your MySQL connection string
	db, err = gorm.Open("mysql", "atul:Choice123@tcp(localhost:3306)/leadapi")
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return
	}

	// Test the connection
	err = db.DB().Ping()
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return
	}

	fmt.Println("Database connection successful")
}

// User model
type User struct {
	gorm.Model
	Username string
	Email    string
}

// Handler to create a new user
func createUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&newUser)
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": newUser})
}

// Handler to fetch all users
func getUsers(c *gin.Context) {
	var users []User
	db.Find(&users)
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func main() {
	defer db.Close()

	// Set up Gin router
	r := gin.Default()

	// Define routes
	r.POST("/users", createUser)
	r.GET("/users", getUsers)

	// Run the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
