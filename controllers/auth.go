package controllers

import (
	"net/http"
	"os"
	"test-golang/config/mysql"
	"test-golang/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var dataRegister struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

var dataLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	// get data
	if c.Bind(&dataRegister) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "data failed"})
		return
	}
	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(dataRegister.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "hash password failed"})
		return
	}

	// create user
	user := models.User{Email: dataRegister.Email, Password: string(hash), Name: dataRegister.Name}
	result := mysql.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user create failend"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "user registered"})
}

func Login(c *gin.Context) {
	// inisialisasi model user
	var user models.User

	// bind data
	if c.Bind(&dataLogin) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "data failed"})
		return
	}

	// query berdasarkan email
	mysql.DB.First(&user, "email = ?", dataLogin.Email)

	// cek apakah user ada
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	// cek password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dataLogin.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password not found"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		// 2 jam token
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "TOKEN GAGAL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
