package controller

import (
	"github.com/amaan287/flightApiGo/initilizers"
	"github.com/amaan287/flightApiGo/models"
	"github.com/gin-gonic/gin"
)

func init() {
	initilizers.LoadEnv()
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
