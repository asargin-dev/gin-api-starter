package user

import (
	"fmt"
	"github.com/asargin-dev/soundproof-backend-demo/internal/models"
	"github.com/asargin-dev/soundproof-backend-demo/pkg/tokenizer"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserProfile(c *gin.Context) {
	var userProfile models.User

	userId, err := tokenizer.ExtractTokenID(c)
	if err != nil {
		return
	}

	profile, err := userProfile.GetUserProfile(userId)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{"userProfile": profile})
}

func UpdateUserProfile(c *gin.Context) {

	var userProfile models.User
	var meta models.Metamask

	if err := c.ShouldBindJSON(&meta); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	publicAddress := meta.ExtractPublicAddress()
	stringTypePublicAddress := publicAddress.String()

	userId, err := tokenizer.ExtractTokenID(c)
	if err != nil {
		return
	}

	profile, err := userProfile.GetUserProfile(userId)
	if err != nil {
		return
	}

	updatedUserProfile, err := profile.UpdateUserProfile(stringTypePublicAddress)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, gin.H{
		"message":     "User profile successfully updated ",
		"userProfile": updatedUserProfile,
	})
}
