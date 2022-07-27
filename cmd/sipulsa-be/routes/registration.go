package routes

import (
	"sipulsa-be/controllers"

	"github.com/labstack/echo/v4"
)

func initUserRoutes(baseApi string, e *echo.Echo) {
	var baseUserApi = baseApi + "/user"

	e.POST(baseUserApi+"/register", controllers.AddNewUserTemp)
	e.GET(baseUserApi+"/all", controllers.FindAllUserTemp)
	// e.POST(baseApi + "/validate")
}
