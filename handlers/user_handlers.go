package handlers

import (
	"github.com/aliparlakci/armut-backend-assessment/models"
	"github.com/aliparlakci/armut-backend-assessment/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Signup(usersCreator services.UserCreator, userGetter services.UserGetter) gin.HandlerFunc {
	return func(c *gin.Context) {
		var creds models.AuthForm
		if err := c.Bind(&creds); err != nil {
			c.String(http.StatusBadRequest, "")
			return
		}

		exists, err := userGetter.UserExists(c.Copy(), creds.Username)
		if err != nil {
			c.String(http.StatusInternalServerError, "")
			return
		}
		if exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
			return
		}

		if err := usersCreator.CreateUser(c.Copy(), creds.Username, creds.Password); err != nil {
			c.String(http.StatusInternalServerError, "")
			return
		}

		c.JSON(http.StatusCreated, gin.H{"result": "user created"})
	}
}