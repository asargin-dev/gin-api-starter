package auth

import (
	"github.com/asargin-dev/soundproof-backend-demo/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	// register if there is not

	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	result := newUser.CheckIfRegistered()

	if result {

		c.JSON(409, gin.H{
			"message": "User already has been registered",
		})

	} else {
		_, err := newUser.Register()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		c.JSON(200, gin.H{
			"message": "Successfully registered",
		})
	}

}

func Login(c *gin.Context) {

	var loginUser models.User

	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	token, err := loginUser.Login()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email or Password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": loginUser, "token": token})
}
