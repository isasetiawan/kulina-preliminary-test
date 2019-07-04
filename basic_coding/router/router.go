package router

import (
	"hello/basic_coding/controller"

	"gopkg.in/go-playground/validator.v9"

	"github.com/labstack/echo"
)

type ReviewValidator struct {
	validator *validator.Validate
}

func (rv *ReviewValidator) Validate(i interface{}) error {
	return rv.validator.Struct(i)
}

func Init() *echo.Echo {
	e := echo.New()
	e.Validator = &ReviewValidator{validator: validator.New()}
	e.POST("/ratings/", controller.CreateRating)
	e.GET("/ratings/", controller.GetRating)
	e.GET("/ratings/:id", controller.GetRatingById)
	e.PUT("/ratings/:id", controller.UpdateRating)
	e.DELETE("/ratings/:id", controller.DeleteRating)

	return e
}
