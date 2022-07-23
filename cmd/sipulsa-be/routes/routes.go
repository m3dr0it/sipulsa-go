package routes

import (
	"net/http"
	"sipulsa-be/controllers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(ctx echo.Context) error {
		ctx.String(http.StatusOK, "Tes")
		return nil
	})

	e.GET("/users", controllers.FindAllUserTemp)
	e.POST("/user", controllers.AddNewUserTemp)

	return e
}
