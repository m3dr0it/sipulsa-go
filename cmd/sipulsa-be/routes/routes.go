package routes

import (
	"net/http"
	"sipulsa-be/controllers"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(ctx echo.Context) error {
		ctx.String(http.StatusOK, "Tes")
		return nil
	})

	e.GET("/users", controllers.FindAllUserTemp)

	return e
}
