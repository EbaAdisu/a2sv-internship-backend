package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// ...

// Global variable to store users in-memory
var users = make(map[string]*User)
var jwtSecret = []byte("your_jwt_secret")

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Go Authentication and Authorization tutorial!",
		})
	})

	router.POST("/register", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request payload"})
			return
		}

		// TODO: Implement user registration logic
		// User registration logic
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(500, gin.H{"error": "Internal server error"})
			return
		}

		user.Password = string(hashedPassword)
		users[user.Email] = &user

		c.JSON(200, gin.H{"message": "User registered successfully"})
	})

	router.POST("/login", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request payload"})
			return
		}
		existingUser, ok := users[user.Email]
		if !ok || bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)) != nil {
			c.JSON(401, gin.H{"error": "Invalid email or password"})
			return
		}

		// TODO: Implement user login logic
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": existingUser.ID,
			"email":   existingUser.Email,
		})

		jwtToken, err := token.SignedString(jwtSecret)
		if err != nil {
			c.JSON(500, gin.H{"error": "Internal server error"})
			return
		}

c.JSON(200, gin.H{"message": "User logged in successfully", "token": jwtToken})
	})

	router.Run()
}
