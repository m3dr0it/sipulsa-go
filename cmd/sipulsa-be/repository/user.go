package repository

import (
	"log"
	"sipulsa-be/database"
	"sipulsa-be/models"
)

func FindAllMember() []models.UserTemp {
	pgcon := database.Connection()
	trx, err := pgcon.Begin()

	defer trx.Close()
	if err != nil {
		log.Println(err.Error())
	}

	var userTemps []models.UserTemp
	errQ := trx.Model(userTemps).ColumnExpr("*").Select()

	if errQ != nil {
		log.Println(errQ.Error())
	}

	log.Println(userTemps)

	log.Println(pgcon)

	return userTemps
}
