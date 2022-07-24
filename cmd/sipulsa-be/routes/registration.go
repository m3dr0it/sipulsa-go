package routes

import (
	"sipulsa-be/controllers"

	"github.com/labstack/echo/v4"
)

func initUserRoutes(baseApi string, e *echo.Echo) {
	e.POST(baseApi+"/register", controllers.AddNewUserTemp)
	// e.POST(baseApi + "/validate")
}
