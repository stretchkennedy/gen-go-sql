package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchkennedy/go-gen-sql/example/data"
)

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}
	r := data.NewFooRepo(db)
	err = r.Create(&data.Foo{})
	if err != nil {
		panic(err)
	}
}
