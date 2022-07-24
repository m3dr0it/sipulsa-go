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

	pgcon.Where(&models.UserTemp{
		Username: username, IsDeleted: false,
	}).
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

	pgcon.Where(&models.UserTemp{PhoneNumber: phoneNumber,
		IsDeleted: false}).
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

	pgcon.Where(&models.UserTemp{Email: mail, IsDeleted: false}).
		Find(&userTemp).Count(&isUserTempExists)

	if isUserTempExists > 0 {
		return true, errors.New(mail + " exists")
	} else {
		return false, nil
	}
}
func IsExistsByUsernameEmailPhoneOtp(username string,
	email string, phone string, otp string) (models.UserTemp, error) {

	var isUserTempExists int64

	var userOtp models.UserTemp

	pgcon.Where(&models.UserTemp{
		Username:    username,
		Email:       email,
		PhoneNumber: phone,
		Otp:         otp,
		IsDeleted:   false,
	}).Find(&userOtp).Count(&isUserTempExists)

	if isUserTempExists > 0 {
		return userOtp, nil
	} else {
		return userOtp, errors.New("Kode Otp Salah")
	}

}
