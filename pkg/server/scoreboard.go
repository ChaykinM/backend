package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) getTasksCount(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "The user is not logged in to the system",
		})
		return
	}
	tasks_stats, err := s.database.GetTasksCount()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"tasks_stats": tasks_stats,
	})
}

func (s *Server) getUserRating(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "The user is not logged in to the system",
		})
		return
	}

	auth_userID, _ := c.Get("UserID")

	user_id, err := strconv.Atoi(c.Param("user_id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userStatus, _ := c.Get("Status")

	if auth_userID == user_id || userStatus != "employee" {
		user_rating, err := s.database.GetUserRating(user_id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"user_rating": user_rating,
		})
	} else {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Insufficient rights to execute the request",
		})
	}
}

func (s *Server) getRatings(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "The user is not logged in to the system",
		})
		return
	}
	ratings, err := s.database.GetRatings()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ratings": ratings,
	})
}
