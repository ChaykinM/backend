package server

import (
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"main.go/pkg/models"
	// "../models"
)

const (
	auth_signingKey = "6ZQRqJ1hbMyNRshyvWBhkQuSDEFoOs3b"
)

func (s *Server) generateAuthToken(authData models.AuthUserData) string {
	authLifetime := 12 // 12 часов срок жизни токена авторизации
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.AuthTokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(authLifetime) * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		authData,
	})
	jwtTokenStr, _ := jwtToken.SignedString([]byte(auth_signingKey))
	return jwtTokenStr
}

func (s *Server) parseAuthToken(jwtToken string) (*models.AuthTokenClaims, error) {
	var claims *models.AuthTokenClaims = new(models.AuthTokenClaims)
	token, err := jwt.ParseWithClaims(jwtToken, &models.AuthTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(auth_signingKey), nil
	})
	if err != nil {
		return claims, err
	}

	claims, _ = token.Claims.(*models.AuthTokenClaims)

	return claims, nil
}

func (s *Server) userIdentification(c *gin.Context) {
	var userAuthorized bool
	var userID int
	var status string
	authToken := c.GetHeader("Authorization")
	if authToken != "" {
		splitToken := strings.Split(authToken, "Bearer ")
		reqToken := splitToken[1]
		if authClaims, authClaims_err := s.parseAuthToken(reqToken); authClaims_err == nil {
			userID = authClaims.UserID
			userAuthorized = true
			status = authClaims.Status
		}
	}

	c.Set("UserAuthorized", userAuthorized)
	c.Set("UserID", userID)
	c.Set("Status", status)
}
