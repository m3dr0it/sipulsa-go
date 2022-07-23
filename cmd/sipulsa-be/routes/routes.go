package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	var apiVersion = "v1"
	var baseApi = "/api/" + apiVersion

	e := echo.New()

	e.GET("/", func(ctx echo.Context) error {
		ctx.String(http.StatusOK, "Tes")
		return nil
	})

	initUserRoutes(baseApi, e)

	return e
}
