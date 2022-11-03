package api

import (
	"advance-test-survey/db"
	"advance-test-survey/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetSurveys(c echo.Context) error {
	db := db.DbManager()
	surveys := []model.Survey{}
	db.Find(&surveys)
	return c.JSON(http.StatusOK, surveys)
}

func CreateSurvey(c echo.Context) error {
	db := db.DbManager()
	s := new(model.Survey)
	if err := c.Bind(s); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	result := db.Create(&s)
	return c.JSON(http.StatusOK, result.Value)
}

func UpdateSurvey(c echo.Context) error {
	db := db.DbManager()
	s := new(model.Survey)
	id := c.Param("id")
	if err := c.Bind(s); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	result := db.Model(&s).Where("id = ?", id).Update(&s)
	return c.JSON(http.StatusOK, result.Value)
}

func DeleteSurvey(c echo.Context) error {
	db := db.DbManager()
	survey := model.Survey{}
	id := c.Param("id")
	db.Where("id = ?", id).Delete(&survey)
	return c.JSON(http.StatusOK, "Deleted")
}

func GetSurvey(c echo.Context) error {
	db := db.DbManager()
	survey := model.Survey{}
	id := c.Param("id")
	db.Where("id = ?", id).First(&survey)
	if survey.ID == 0 {
		return c.JSON(http.StatusNotFound, "Data not found")
	}
	return c.JSON(http.StatusOK, survey)
}
