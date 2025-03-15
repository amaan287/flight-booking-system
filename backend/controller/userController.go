package controller

import (
	"time"

	"github.com/amaan287/flightApiGo/initilizers"
	"github.com/amaan287/flightApiGo/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func HashedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

var secret1 = []byte("Secret")

func generateJWT(UserID int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = UserID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString(secret1)
	if err != nil {
		return "Error signing jwt token", err
	}
	return tokenString, nil
}

func Signup(c *gin.Context) {
	//body, _ := io.ReadAll(c.Request.Body)
	//	fmt.Println(string(body))
	var User struct {
		Name     string
		Email    string
		Password string
		Phone    string
	}
	if err := c.Bind(&User); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request", "Error": err,
		})
		return
	}
	hashedPassword, err := HashedPassword(User.Password)
	if err != nil {
		c.JSON(400, gin.H{
			"Message": "Error hashing password",
		})
		return
	}
	user := models.User{
		Name:     User.Name,
		Email:    User.Email,
		Phone:    User.Phone,
		Password: hashedPassword}
	userRes := initilizers.DB.Create(&user)
	if userRes.Error != nil {
		c.JSON(400, gin.H{
			"message": "Error creating user", "error": userRes.Error,
		})
		return
	}
	token, signError := generateJWT(user.ID)
	if signError != nil {
		c.JSON(400, gin.H{
			"message": "Error generating jwt token",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "User signed up sucessfully",
		"data": models.AuthResponse{
			Token: token,
			User: models.User{
				Name:     User.Name,
				Email:    User.Email,
				Phone:    User.Phone,
				Password: hashedPassword,
			}},
	})

}
