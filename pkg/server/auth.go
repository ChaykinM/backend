package server

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/pkg/models"
)

func (s *Server) loginHandler(c *gin.Context) {
	var loginRequest models.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errors.New("Incorrect authorization data").Error()})
	} else {
		passwordHash := md5.New()
		passwordHash.Write([]byte(loginRequest.Password))
		passwordMd5 := hex.EncodeToString(passwordHash.Sum(nil))
		loginRequest.Password = passwordMd5

		if authData, err := s.database.LoginAuthorization(&loginRequest); err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": errors.New("Incorrect username or password. Please check your login details.").Error()})
		} else {
			authToken := s.generateAuthToken(authData)
			c.JSON(http.StatusOK, gin.H{
				"Message": "Successful authorization",
				"Token":   authToken,
			})
		}
	}
}

func (s *Server) registerHandler(c *gin.Context) {
	var registerRequest models.RegisterRequest
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errors.New("Incorrect data for registration").Error()})
	} else {
		passwordHash := md5.New()
		passwordHash.Write([]byte(registerRequest.Password))
		passwordMd5 := hex.EncodeToString(passwordHash.Sum(nil))
		registerRequest.Password = passwordMd5
		if authData, err := s.database.RegisterUser(&registerRequest); err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": errors.New("Such a user already exists. Change the registration data.").Error()})
		} else {
			authToken := s.generateAuthToken(authData)

			c.JSON(http.StatusOK, gin.H{
				"Message": "Successful registration",
				"Token":   authToken,
			})
		}
	}
}
