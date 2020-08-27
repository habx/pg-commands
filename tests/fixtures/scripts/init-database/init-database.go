package main

import (
	"io/ioutil"
	"os"
	"strconv"

	"github.com/go-pg/pg"
	"github.com/habx/pg-commands/tests/fixtures"
)

func main() {
	config := fixtures.Setup()
	querys, err := ioutil.ReadFile("tests/fixtures/scripts/01_database.sql")
	if err != nil {
		panic(err)
	}

	db := pg.Connect(&pg.Options{
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DB"),
		Addr:     config.Host + ":" + strconv.Itoa(config.Port),
	})
	defer db.Close()
	_, err = db.Exec(`CREATE ROLE "dev_example" WITH LOGIN ENCRYPTED PASSWORD 'password';`)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(`CREATE DATABASE "dev_example" OWNER "dev_example";`)
	if err != nil {
		panic(err)
	}
	db = pg.Connect(&pg.Options{
		User:     config.Username,
		Password: config.Password,
		Database: config.DB,
		Addr:     config.Host + ":" + strconv.Itoa(config.Port),
	})
	defer db.Close()
	_, err = db.Exec(string(querys))
	if err != nil {
		panic(err)
	}
}
