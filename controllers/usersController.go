package controllers

import (
	"fmt"
	"github.com/bugsssssss/auth-gin/initizalizers"
	"github.com/bugsssssss/auth-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"time"
)

func Signup(c *gin.Context) {
	//	get email and password from req body
	var body struct {
		Email    string
		Password string
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read or validate the request body",
		})
		return
	}

	email := body.Email
	password := body.Password

	if email == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Either email or password wasn't provided",
		})
		return
	}

	log.Printf("\nEmail: %s\nPassword: %s", email, password)

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	user := models.User{Email: body.Email, Password: string(hash)}
	result := initizalizers.DB.Create(&user) // pass pointer of data to Create

	if result.Error != nil {
		log.Printf("Error: %s", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("failed to create user with error - %s", result.Error),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user created successfully",
	})
}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}
	user := models.User{}
	initizalizers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(400, gin.H{
			"error": "invalid email",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid password",
		})
		return
	}

	// creating token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		log.Printf("Error while creating a token: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to generate token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func Validate(c *gin.Context) {

	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})

}
