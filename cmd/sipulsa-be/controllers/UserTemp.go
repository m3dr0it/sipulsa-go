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

	if err := c.Bind(ut); err != nil {
		return echo.ErrInternalServerError
	}

	if err := c.Validate(ut); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	errService := registration.Register(ut)

	if errService != nil {
		c.JSON(http.StatusConflict, responses.ConflictResponse(errService))
	} else {
		c.JSON(http.StatusCreated, responses.CreatedResponse("otp send to email"))
	}

	return nil
}

func ValidateOtp(c echo.Context) error {
	ut := new(requests.NewUserTemp)
	err := c.Bind(ut)

	if err != nil {
		log.Println(err.Error())
	}

	return nil
}
