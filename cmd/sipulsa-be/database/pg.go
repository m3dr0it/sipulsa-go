package database

import (
	"context"
	"log"

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

	ctx := context.Background()

	if err := pgCon.Ping(ctx); err != nil {
		panic(err)
	}

	log.Println("connected to database")

	return pgCon

}
