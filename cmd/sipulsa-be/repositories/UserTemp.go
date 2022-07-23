package repositories

import (
	"log"
	database "sipulsa-be/databases"
	"sipulsa-be/models"
	"sipulsa-be/models/requests"
	"sipulsa-be/utilities/otp"
	"time"

	"github.com/go-pg/pg"
)

var pgcon = new(pg.DB)

func init() {
	pgcon = database.Connection()
}

func FindAllUserTemps() []models.UserTemp {
	trx, err := pgcon.Begin()
	defer trx.Close()

	if err != nil {
		log.Println(err.Error())
	}

	var userTemps []models.UserTemp
	errQ := trx.Model(&userTemps).ColumnExpr("*").Select()

	if errQ != nil {
		log.Println(errQ.Error())
	}

	log.Println(userTemps)

	log.Println(pgcon)

	return userTemps
}

func AddUserTemp(ut requests.NewUserTemp) {
	trx, err := pgcon.Begin()
	defer trx.Close()

	if err != nil {
		log.Println(err.Error())
	}

	userTemp := &models.UserTemp{
		Name:                     ut.Name,
		Username:                 ut.Username,
		Email:                    ut.Email,
		Otp:                      otp.GenerateOtp(),
		OtpExpired:               time.Now().Add(time.Duration(time.Minute.Minutes() * 5)),
		CreatedAt:                time.Now(),
		RegisteredByReferralCode: ut.ReferralCode,
		PhoneNumber:              ut.PhoneNumber,
		CreatedBy:                "System",
	}

	insert, errInsert := trx.Model(userTemp).Insert()

	log.Println(insert)

	if errInsert != nil {
		log.Println(errInsert)
	}
}
