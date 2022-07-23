package controllers

import (
	"net/http"
	"sipulsa-be/models/requests"
	models "sipulsa-be/models/responses"
	"sipulsa-be/repositories"
	"time"

	"github.com/labstack/echo/v4"
)

func FindAllUserTemp(ec echo.Context) error {
	usersTemp := repositories.FindAllUserTemps()

	res := models.Response{
		Timestamp:  time.Now(),
		Message:    "null",
		StatusCode: http.StatusFound,
		Content:    usersTemp,
	}

	err := ec.JSON(http.StatusOK, res)

	if err != nil {
		return err
	}

	return nil
}

func addUserTemp(c echo.Context) error {
	ut := new(requests.NewUserTemp)

	if err := c.Bind(ut); err != nil {
		return err
	}

	return nil
}
