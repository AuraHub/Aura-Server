package controllers

import (
	"Aura-Server/initializers"
	"Aura-Server/models"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// Get the credentials off req body
	var body struct {
		Name     string
		LastName string
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// Create the user
	user := models.User{
		Name:      body.Name,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  string(hash),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := initializers.Database.Collection("users").InsertOne(context.TODO(), user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
			"user":  user,
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

func Login(c *gin.Context) {
	// Get the email/pass off req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Look up requested user
	filter := bson.D{{Key: "email", Value: body.Email}}
	var user models.User
	err := initializers.Database.Collection("users").FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	if user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Compare sent in pass with saved user pass hash
	err1 := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	// send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged in",
		"data":    user,
	})
}

func Validate(c *gin.Context) {
	var user, _ = c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"data":    user,
		"message": "You are logged in",
	})
}

func Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged out",
	})
}

func GetUser(c *gin.Context) {
	var user, _ = c.Get("user")

	c.JSON(http.StatusOK, user)
}
