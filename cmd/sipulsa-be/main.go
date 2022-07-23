package main

import (
	database "sipulsa-be/databases"
	"sipulsa-be/routes"
)

func main() {
	e := routes.Init()

	err := database.IsConnectedToDb()
	if err != nil {
		panic("database not connected")
	}

	e.Start(":8089")
}
