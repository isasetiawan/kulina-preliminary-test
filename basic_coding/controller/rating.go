package controller

import (
	"fmt"
	"hello/basic_coding/database"
	"hello/basic_coding/model"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func CreateRating(c echo.Context) error {
	review := new(model.UserReview)
	err := c.Bind(review)
	err = c.Validate(review)
	if err != nil {
		response := strings.Split(fmt.Sprintln(err), "\n")
		return c.JSONPretty(http.StatusBadRequest, map[string]interface{}{"message": response}, "	")
	}

	db := database.DbManager()
	db.Create(&review)

	return c.JSON(http.StatusCreated, review)
}

func GetRating(c echo.Context) error {
	db := database.DbManager()
	rating := []model.UserReview{}
	db.Find(&rating)

	return c.JSON(http.StatusOK, rating)
}

func GetRatingById(c echo.Context) error {
	db := database.DbManager()
	id := c.Param("id")
	review := new(model.UserReview)
	if err := db.Where("id = ?", id).First(&review).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "notfound"})
	}

	return c.JSON(http.StatusOK, review)
}

func UpdateRating(c echo.Context) error {
	db := database.DbManager()
	review := new(model.UserReview)
	id := c.Param("id")
	err := c.Bind(review)
	err = c.Validate(review)
	if err != nil {
		response := strings.Split(fmt.Sprintln(err), "\n")
		return c.JSONPretty(http.StatusBadRequest, map[string]interface{}{"message": response}, "	")
	}
	old_review := new(model.UserReview)

	if err := db.Where("id = ?", id).First(&old_review).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "notfound"})
	}
	db.Model(&old_review).Updates(review)
	return c.JSON(http.StatusCreated, review)

}

func DeleteRating(c echo.Context) error {
	db := database.DbManager()
	id := c.Param("id")
	review := new(model.UserReview)
	if err := db.Where("id = ?", id).First(&review).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "notfound"})
	}
	db.Delete(&review)
	return c.JSON(http.StatusOK, review)

}
