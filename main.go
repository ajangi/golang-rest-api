package main

import (
	"net/http"

	"github.com/ajangi/golang-rest-api/utils"
	"github.com/ajangi/golang-rest-api/controllers"
	"github.com/labstack/echo"
)

func main() {
	echo.NotFoundHandler = func(c echo.Context) error {
		emptyData := utils.ResponseData{}
		notFoundMessage := utils.ResponseMessages{utils.GetMessageByKey(utils.NotFoundErrorMessageKey)}
		errorResponse := utils.ResponseApi{Result: "ERROR", Data: emptyData, Messages: notFoundMessage, StatusCode: http.StatusNotFound}
		return c.JSON(http.StatusNotFound, errorResponse)
	}
	e := echo.New()
	usersGroup := e.Group("/users")
	usersGroup.POST("", controllers.RegisterUser)
	e.Logger.Fatal(e.Start(":1323"))
}
