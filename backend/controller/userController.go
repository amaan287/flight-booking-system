package controller

import (
	"fmt"
	"os"
	"time"

	"github.com/amaan287/flightApiGo/initilizers"
	"github.com/amaan287/flightApiGo/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	initilizers.LoadEnv()
}
func HashedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

var secret1 []byte = []byte(os.Getenv("JWT_SECRET"))

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
	var Body struct {
		Name     string
		Email    string
		Password string
		Phone    string
	}
	var User models.User
	if err := c.Bind(&Body); err != nil {
		c.JSON(400, gin.H{
			"data": models.ErrorResponse{
				Error:   err.Error(),
				Message: "Invalid request",
			},
		})
		return
	}
	initilizers.DB.Where("email = ?", Body.Email).First(&User)
	if User.Email != "" {
		c.JSON(400, gin.H{
			"data": models.ErrorResponse{
				Error:   " ",
				Message: "User with this email already Exist",
			}})
		return
	}

	hashedPassword, err := HashedPassword(Body.Password)
	if err != nil {
		c.JSON(400, gin.H{
			"Message": "Error hashing password",
		})
		return
	}
	user := models.User{
		Name:     Body.Name,
		Email:    Body.Email,
		Phone:    Body.Phone,
		Password: hashedPassword}
	userRes := initilizers.DB.Create(&user)
	if userRes.Error != nil {
		c.JSON(400, gin.H{
			"message": "Error creating user", "error": userRes.Error,
		})
		return
	}
	fmt.Println(user.ID)
	token, signError := generateJWT(user.ID)

	if signError != nil {
		c.JSON(400, gin.H{
			"data": models.ErrorResponse{
				Error:   signError.Error(),
				Message: "Error signing jwt",
			}})
		return
	}
	c.JSON(200, gin.H{
		"message": "User signed up sucessfully",
		"data": models.AuthResponse{
			Token: token,
			User: models.User{
				Name:  Body.Name,
				Email: Body.Email,
				Phone: Body.Phone,
			}},
	})
}

func Signin(c *gin.Context) {
	var Body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var User models.User
	err := c.Bind(&Body)
	if err != nil {

		return
	}
	fmt.Println(Body.Password, Body.Email)
	res := initilizers.DB.Where("email = ?", Body.Email).First(&User)
	if res.Error != nil {
		c.JSON(400, gin.H{
			"data": models.ErrorResponse{
				Error:   res.Error.Error(),
				Message: "User not found",
			}})
		return
	}
	match := CheckPasswordHash(Body.Password, User.Password)
	if !match {
		c.JSON(400, gin.H{
			"data": models.ErrorResponse{
				Error:   " ",
				Message: "Wrong password",
			}})
		return
	}
	fmt.Println(User.ID)

	token, jwtError := generateJWT(User.ID)
	if jwtError != nil {
		c.JSON(400, gin.H{
			"data": models.ErrorResponse{
				Message: "Error generating jwt token",
				Error:   jwtError.Error(),
			}})
	}
	c.JSON(200, gin.H{
		"message": "Signin Success",
		"data": models.AuthResponse{
			Token: token,
			User:  User,
		}})
}
func UpdatePass(c *gin.Context) {
	// get user information using userId
	id := c.Param("id")
	var user models.User
	initilizers.DB.Where("id=?", id).First(&user)

	// Create a struct with both password fields
	var passwords struct {
		OldPassword string `json:"oldPassword"`
		NewPassword string `json:"newPassword"`
	}

	// Bind JSON data to the struct (once)
	if err := c.ShouldBindJSON(&passwords); err != nil {
		c.JSON(400, gin.H{
			"data": models.ErrorResponse{
				Error:   err.Error(),
				Message: "Error binding password data",
			},
		})
		return
	}

	// Check old password
	match := CheckPasswordHash(passwords.OldPassword, user.Password)
	if !match {
		c.JSON(400, gin.H{
			"data": models.ErrorResponse{
				Message: "Incorrect password",
			},
		})
		return
	}

	// Hash and update new password
	hashedPassword, err := HashedPassword(passwords.NewPassword)
	if err != nil {
		c.JSON(400, gin.H{
			"data": models.ErrorResponse{
				Message: "Error hashing password",
				Error:   err.Error(),
			},
		})
		return
	}

	// Update user password in database
	result := initilizers.DB.Model(&user).Update("password", hashedPassword)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"data": models.ErrorResponse{
				Message: "Failed to update in database",
				Error:   result.Error.Error(),
			}})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(500, gin.H{
			"data": models.ErrorResponse{
				Message: "No rows were updated",
			},
		})
	}
	c.JSON(200, gin.H{
		"message": "Password changed successfully",
	})
}

func UpdateName(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	initilizers.DB.Where("id=?", id).First(&user)
	var Name struct {
		OldName  string `json:"oldName"`
		NewName  string `json:"newName"`
		Password string `json:"password"`
	}
	if err := c.ShouldBind(&Name); err != nil {
		c.JSON(400, gin.H{
			"data": models.ErrorResponse{
				Error:   err.Error(),
				Message: "Error getting the name",
			}})
		return
	}
	//check users password
	match := CheckPasswordHash(Name.Password, user.Password)
	if !match {
		c.JSON(400, gin.H{
			"data": models.ErrorResponse{
				Message: "Incorrect password",
			}})
		return
	}
	result := initilizers.DB.Model(&user).Update("name", Name.NewName)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"data": models.ErrorResponse{
				Error:   result.Error.Error(),
				Message: "Error updating the name in database",
			}})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(500, gin.H{
			"data": models.ErrorResponse{
				Message: "no rows were updated",
			}})
		return
	}
	c.JSON(200, gin.H{
		"message": "name updated successfully",
		"User":    user,
	})

}
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	initilizers.DB.Where("id=?", id).First(&user)
	c.JSON(200, gin.H{
		"user": user,
	})
}
