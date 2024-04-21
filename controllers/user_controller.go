package controllers

import (
	data "go-side-project/initializers"
	model "go-side-project/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)


func Login(c *gin.Context){
	var user model.User

	if err := c.ShouldBind(&user);err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request data"})
		return
	}

	var foundUser model.User
	if err := data.DB.Where("email = ?", user.Email).First(&foundUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect email"})
		return
	}

    user.Password = user.Password + "safety"
	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
		return
	}

    // Using JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "ID":    foundUser.ID,
        "name":  foundUser.Name,
        "email": foundUser.Email,
        "role":  foundUser.RoleID,
        "exp":   time.Now().Add(time.Hour * 24 * 30).Unix(), // 设置过期时间为 30 天
    })


    // Salted
	secretKey := os.Getenv("SECRET_KEY")
	log.Println("secretKey", secretKey)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}
	c.SetSameSite(http.SameSiteDefaultMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", true, true)
    // G 傳送 response
	c.JSON(http.StatusOK, gin.H{"token": tokenString, "message": "User login successfully"})

}

func Logout(c *gin.Context){
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}

func Register(c *gin.Context){
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request data"})
		return
	}
	user.Password = user.Password + "safety"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encrypt password"})
        return
    }

	user.Password = string(hashedPassword)

	if err := data.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}