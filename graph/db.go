package graph

import (
	"os"

	"github.com/go-pg/pg"
)

func Connect() *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: "graphql",
	})

	return db
}
