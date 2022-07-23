package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {

	dsn := "host=localhost user=root password=H8dVeZYjv66xfXbq dbname=sipulsa port=5433 sslmode=disable TimeZone=Asia/Shanghai"

	pgCon, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return pgCon

}

func IsConnectedToDb() error {
	con := Connection()
	db, err := con.DB()

	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		return err
	}

	return nil
}
