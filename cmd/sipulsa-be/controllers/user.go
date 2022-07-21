package controllers

import (
	"net/http"
	models "sipulsa-be/models/responses"
	"sipulsa-be/repositories"
	"time"

	"github.com/labstack/echo"
)

func FindAllUserTemp(ec echo.Context) error {
	usersTemp := repositories.FindAllMember()

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
