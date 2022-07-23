package repositories

import (
	"database/sql"
	"errors"
	"log"
	database "sipulsa-be/databases"
	models "sipulsa-be/models/pg"

	"gorm.io/gorm"
)

var pgcon = new(gorm.DB)
var db = &sql.DB{}

func init() {
	pgcon = database.Connection()

}

func FindAllUserTemps() []models.UserTemp {
	var userTemps []models.UserTemp

	result := pgcon.Find(&userTemps)
	if result.Error != nil {
		log.Println(result.Error)
	}

	return userTemps
}

func AddUserTemp(userTemp models.UserTemp) error {

	result := pgcon.Create(&userTemp)

	if result != nil {
		log.Println(result)
		return result.Error
	}

	return nil
}

func IsExistsByUsername(username string) (bool, error) {
	var isUserTempExists int64

	var userTemp models.UserTemp

	pgcon.Where(&models.UserTemp{Username: username}).
		Find(&userTemp).Count(&isUserTempExists)

	if isUserTempExists > 0 {
		return true, errors.New(username + " exists")
	} else {
		return false, nil
	}
}

func IsExistsByPhoneNumber(phoneNumber string) (bool, error) {
	var isUserTempExists int64

	var userTemp models.UserTemp

	pgcon.Where(&models.UserTemp{PhoneNumber: phoneNumber}).
		Find(&userTemp).Count(&isUserTempExists)

	if isUserTempExists > 0 {
		return true, errors.New(phoneNumber + " exists")
	} else {
		return false, nil
	}
}

func IsExistsByEmail(mail string) (bool, error) {
	var isUserTempExists int64

	var userTemp models.UserTemp

	pgcon.Where(&models.UserTemp{Email: mail}).
		Find(&userTemp).Count(&isUserTempExists)

	if isUserTempExists > 0 {
		return true, errors.New(mail + " exists")
	} else {
		return false, nil
	}
}
