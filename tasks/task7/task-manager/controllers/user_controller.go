package controllers

import (
    "net/http"
    "task_manager/data"
    "task_manager/models"
    "time"

    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// JWT secret key (in production, store this securely)
var jwtSecret = []byte("your_secret_key")

// Claims defines the structure of JWT claims.
type Claims struct {
    UserID   primitive.ObjectID `json:"user_id"`
    Username string             `json:"username"`
    Role     string             `json:"role"`
    jwt.StandardClaims
}

// UserController for user management
type UserController struct {
    UserService *data.UserService
}

func NewUserController(userService *data.UserService) *UserController {
    return &UserController{UserService: userService}
}

// Register handles POST /register
func (uc *UserController) Register(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
        return
    }

    createdUser, err := uc.UserService.RegisterUser(user)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }

    // Do not return password in response
    createdUser.Password = ""
    c.JSON(http.StatusCreated, createdUser)
}

// Login handles POST /login and returns a JWT token upon success.
func (uc *UserController) Login(c *gin.Context) {
    var input struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
        return
    }

    user, err := uc.UserService.LoginUser(input.Username, input.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
        return
    }

    // Create token claims
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        UserID:   user.ID,
        Username: user.Username,
        Role:     user.Role,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }
    // Generate JWT token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}