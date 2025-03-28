package middleware

import (
	"net/http"
	"strings"
	"task_manager/controllers"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware ensures requests include a valid JWT token.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Expect header "Authorization: Bearer <token>"
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization header missing"})
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization header"})
			return
		}

		tokenStr := parts[1]
		claims := &controllers.Claims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return controllers.JwtSecret(), nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			return
		}
		// Store claims in context for access in handlers
		c.Set("user", claims)
		c.Next()
	}
}

// AdminMiddleware checks if the user is an admin.
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}
		claims := user.(*controllers.Claims)
		if claims.Role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Admin access required"})
			return
		}
		c.Next()
	}
}
