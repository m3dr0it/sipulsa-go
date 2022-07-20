package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sipulsa-be/models"
	"sipulsa-be/repository"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", func(ctx echo.Context) error {
		var UserTemp []models.UserTemp

		UserTemp = repository.FindAllMember()

		users, err := json.Marshal(UserTemp)

		if err != nil {
			log.Println(err.Error())
		}

		return ctx.String(http.StatusOK, string(users))
	})

	e.Start(":8089")
}
