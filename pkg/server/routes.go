package server

func (s *Server) routerSettings() {
	s.authRoutes()
	s.userRoutes()
	s.tasksRoutes()
	s.scoreboardRoutes()
}

func (s *Server) authRoutes() {
	authRouters := s.router.Group("/auth")
	{
		authRouters.POST("/login", s.loginHandler)
		authRouters.POST("/register", s.registerHandler)
	}
}

func (s *Server) userRoutes() {
	userRoutes := s.router.Group("users", s.userIdentification)
	{
		userRoutes.GET("/", s.getUsers)
		userRoutes.GET("/:user_id", s.getUserByID)
		userRoutes.POST("/:user_id/edit", s.editUser)
		userRoutes.DELETE("/:user_id/del", s.deleteUser)
		userRoutes.POST("/:user_id/updPass", s.updateUserPassword)
	}
}

func (s *Server) tasksRoutes() {
	tasksRoutes := s.router.Group("/tasks", s.userIdentification)
	{
		tasksRoutes.GET("/", s.getTasks)
		tasksRoutes.GET("/personal_solved/:user_id", s.getUserSolvedTasks)
		tasksRoutes.POST("/:task_id/solve", s.solveTask)
		tasksRoutes.POST("/add", s.addTask)
	}
}

func (s *Server) scoreboardRoutes() {
	scoreboardRoutes := s.router.Group("scoreboard", s.userIdentification)
	{
		scoreboardRoutes.GET("/ratings", s.getRatings)
		scoreboardRoutes.GET("/tasks_count", s.getTasksCount)
		scoreboardRoutes.GET("/personal_rating/:user_id", s.getUserRating)
	}
}
