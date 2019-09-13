package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/phongloihong/go_framework/db/types"
	"github.com/phongloihong/go_framework/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetStudent return all studen in json
func GetStudents(c echo.Context) error {
	students, err := repository.Fetch()
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			types.ErrorResponse{Code: "BadRequest", Message: err.Error()},
		)
	}

	return c.JSON(http.StatusOK, students)
}

func GetStudent(c echo.Context) error {
	objectId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			types.ErrorResponse{Code: "BadRequest", Message: "Invalid id"},
		)
	}

	student, err := repository.GetById(objectId)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			types.ErrorResponse{Code: "BadRequest", Message: err.Error()},
		)
	}

	return c.JSON(http.StatusOK, student)
}

func SearchStudent(c echo.Context) error {
	var req types.StudentSearchRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			types.ErrorResponse{Code: "BadRequest", Message: err.Error()},
		)
	}

	students, err := repository.Find(req)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			types.ErrorResponse{Code: "BadRequest", Message: err.Error()},
		)
	}

	return c.JSON(http.StatusOK, students)
}

// CheckHeath check server status
func CheckHeath(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
