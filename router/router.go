package router

import (
	"github.com/labstack/echo"
	"github.com/phongloihong/go_framework/handlers"
)

// Public collection of public route
func Public(e *echo.Echo) {
	publicRoute := e.Group("/v1/public")

	publicRoute.GET("/health", handlers.CheckHeath)
	publicRoute.GET("/student", handlers.GetStudents)
	publicRoute.GET("/student/id/:id", handlers.GetStudent)
	publicRoute.PATCH("/student", handlers.SearchStudent)
}

// Staff route
func Staff(e *echo.Echo) {
	staffRoute := e.Group("/v1/staff")

	staffRoute.POST("/student", handlers.AddStudent)
	staffRoute.PATCH("/student", handlers.UpdateStudent)
	staffRoute.DELETE("/student/id/:id", handlers.DeleteStudent)
}
