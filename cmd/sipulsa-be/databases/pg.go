package database

import (
	"context"

	"github.com/go-pg/pg/v10"
)

func Connection() *pg.DB {
	dbOpt := pg.Options{
		Network:         "tcp",
		Addr:            "localhost:5433",
		ApplicationName: "sipulsa-go",
		User:            "root",
		Password:        "H8dVeZYjv66xfXbq",
		Database:        "sipulsa",
	}
	pgCon := pg.Connect(&dbOpt)

	return pgCon

}

func IsConnectedToDb() error {
	db := Connection()

	ctx := context.Background()

	if err := db.Ping(ctx); err != nil {
		return err
	}

	return nil
}
