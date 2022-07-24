package main

import (
	database "sipulsa-be/databases"
	"sipulsa-be/routes"
	localvalidator "sipulsa-be/utilities/validator"
)

func main() {
	e := routes.Init()
	localvalidator.SetValidator(e)
	// responses.InitErrorResponse(e)

	err := database.IsConnectedToDb()
	if err != nil {
		panic("database not connected")
	}

	e.Start(":8089")
}
