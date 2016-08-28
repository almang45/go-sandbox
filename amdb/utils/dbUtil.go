package utils

import (
	"database/sql"
	"fmt"

	"github.com/almang45/go-sandbox/amdb/models"
	"github.com/go-gorp/gorp"
)

const (
	dUser      = "postgres"
	dbPassword = "password"
	dbName     = "postgres"
)

func InitMovieDb() *gorp.DbMap {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		dUser, dbPassword, dbName)
	db, err := sql.Open("postgres", dbInfo)
	CheckError(err, "sql.Open failed!")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbmap.AddTableWithName(models.MovieDetails{}, "movies").SetKeys(true, "Id").SetUniqueTogether("title", "year").AddIndex("imdb_id_index", "Btree", []string{"imdb_id"})

	err = dbmap.CreateTablesIfNotExists()
	CheckError(err, "Create table movies failed!")

	err = dbmap.CreateIndex()
	CheckError(err, "Create index on movies.ImdbId failed!")

	return dbmap
}

func InitUserDb() *gorp.DbMap {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		dUser, dbPassword, dbName)
	db, err := sql.Open("postgres", dbInfo)
	CheckError(err, "sql.Open failed!")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbmap.AddTableWithName(models.User{}, "users").SetKeys(true, "Id").AddIndex("email_index", "Btree", []string{"email"})

	err = dbmap.CreateTablesIfNotExists()
	CheckError(err, "Create table users failed!")

	err = dbmap.CreateIndex()
	CheckError(err, "Create index on users.Email failed!")

	return dbmap
}
