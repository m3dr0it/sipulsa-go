package controllers

import (
	"log"
	"net/http"
	"sipulsa-be/models/requests"
	"sipulsa-be/models/responses"
	"sipulsa-be/repositories"
	"sipulsa-be/services/registration"

	"github.com/labstack/echo/v4"
)

func FindAllUserTemp(ec echo.Context) error {
	usersTemp := repositories.FindAllUserTemps()

	err := ec.JSON(http.StatusOK, responses.OkResponse(usersTemp))

	if err != nil {
		return err
	}

	return nil
}

func AddNewUserTemp(c echo.Context) error {
	ut := new(requests.NewUserTemp)
	err := c.Bind(ut)

	if err != nil {
		log.Println(err.Error())
	}

	errService := registration.Register(c, ut)

	if errService != nil {
		c.JSON(http.StatusConflict, responses.ConflictResponse(errService))
	} else {
		c.JSON(http.StatusCreated, responses.CreatedResponse(nil))

	}

	return nil
}
