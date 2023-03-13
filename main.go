package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

// type User struct {
// 	ID                 int    `json:"id"`
// 	Username           string `json:"name"`
// 	Email              string `json:"email"`
// 	PasswordResetToken string `json:"password_reset_token"`
// }

const dbuser = "root"
const dbpass = "Mys22/03/2023ql"
const dbname = "recordings"

func setUpdatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)
	if err != nil {
		logrus.Fatalf("Failed to connect to database: %v", err)
	}
	return db, nil
}

func main() {
	router := gin.Default()
	router.GET("/user/:email", GetToken)

	if err := router.Run(":8080"); err != nil {
		logrus.Fatalf("Failed to start server: %v", err)
	}
}

func GetToken(c *gin.Context) {
	url := "https://portal.dev01.int.betika.com/site/reset-password?token="
	username := c.Param("email")
	fmt.Println(username)

	// Create database connection
	db, err := setUpdatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	defer db.Close()

	// Prepare SQL statement
	var token string
	err = db.QueryRow("SELECT password_reset_token FROM user WHERE email = ?", username).Scan(&token)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Token not found"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
	}
	fmt.Println(token)

	resetUrl := url + token

	c.JSON(http.StatusOK, resetUrl)
}

// Execute query and retrieve result
// var passwordResetToken string
// err = token.QueryRow(username).Scan(&passwordResetToken)
// if err != nil {
// 	log.Fatal(err)
// }

// fmt.Println("Password reset token:", passwordResetToken)

// return url + passwordResetToken

// func getDSN() string {
// 	dbUser := os.Getenv("root")
// 	dbPass := os.Getenv("Mys22/03/2023ql")
// 	dbName := os.Getenv("user")
// 	return dbUser + ":" + dbPass + "@tcp(127.0.0.1:3306)/" + dbName
// }

// package main

// import (
// 	"database/sql"
// 	"log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type User struct {
// 	Id                   int    `json:"id"`
// 	Username             string `json:"name"`
// 	Email                string `json:"email"`
// 	Password_reset_token string `json:"password_reset_token"`
// }

// const dbuser = "root"
// const dbpass = "Mys22/03/2023ql"
// const dbname = "user"

// func main() {
// 	router := gin.Default()
// 	router.GET("/user/{username}", getUser)

// 	router.Run(":8080")
// }

// func getUser(c *gin.Context) {
// 	username := c.Param("username")
// 	passResetLink := GetToken(username)
// 	c.IndentedJSON(http.StatusOK, passResetLink)

// }

// func GetToken(user string) string {
// 	url := "https://portal.dev01.int.betika.com/site/reset-password?token="
// 	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	var token string
// 	if err := db.QueryRow("SELECT Password_reset_token FROM user WHERE email=" + user + ".com").Scan(&token); err != nil {
// 		log.Fatal(err)
// 	}
// 	return url + token
// }
