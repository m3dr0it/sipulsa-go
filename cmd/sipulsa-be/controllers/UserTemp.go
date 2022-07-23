package controllers

import (
	"log"
	"net/http"
	"sipulsa-be/models/requests"
	"sipulsa-be/models/responses"
	"sipulsa-be/repositories"

	"github.com/labstack/echo/v4"
)

func AddNewUserTemp(c echo.Context) error {
	ut := new(requests.NewUserTemp)
	err := c.Bind(ut)

	if err != nil {
		log.Println(err.Error())
	}

	repositories.AddUserTemp(*ut)

	errC := c.JSON(http.StatusCreated, responses.CreateResponse(nil))

	if errC != nil {
		return errC
	}

	return nil
}
