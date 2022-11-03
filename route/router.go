package route

import (
	"advance-test-survey/api"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/survey", api.GetSurveys)
	e.POST("/survey", api.CreateSurvey)
	e.PUT("/survey/:id", api.UpdateSurvey)
	e.DELETE("/survey/:id", api.DeleteSurvey)
	e.GET("/survey/:id", api.GetSurvey)
	return e
}
