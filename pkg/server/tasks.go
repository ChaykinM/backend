package server

import (
	"errors"
	"net/http"
	"strconv"

	"main.go/pkg/models"

	"github.com/gin-gonic/gin"
)

func (s *Server) getTasks(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "The user is not logged in to the system",
		})
		return
	}
	tasks, err := s.database.GetTasks()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}

func (s *Server) getUserSolvedTasks(c *gin.Context) {
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
		solved_tasks, query_err := s.database.GetUserSolvedTasks(user_id)
		if query_err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": query_err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"solved_tasks": solved_tasks,
		})
	} else {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Insufficient rights to execute the request",
		})
	}
}

func (s *Server) solveTask(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "The user is not logged in to the system",
		})
		return
	}

	var solvedTask models.SolveTaskRequest
	if err := c.ShouldBindJSON(&solvedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("Invalid request").Error(),
		})
		return
	}
	id, err := s.database.SolveTask(&solvedTask)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"message": "The task was successfully solved",
	})
}

func (s *Server) addTask(c *gin.Context) {
	userAuthorized, _ := c.Get("UserAuthorized")
	userStatus, _ := c.Get("Status")
	if userAuthorized == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "The user is not logged in to the system",
		})
		return
	}
	if userStatus == "employee" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Insufficient permissions to perform actions",
		})
		return
	}

	var task models.AddTaskRequest
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("Invalid request").Error(),
		})
		return
	}
	id, err := s.database.AddTask(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"message": "The task was successfully added",
	})
}
