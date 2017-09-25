package main

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

// Time is time
type Time struct {
	pq.NullTime
}

var db *sql.DB

// Connect connects to postgres database using datasouce
func Connect(dataSource string) error {
	d, err := sql.Open("postgres", dataSource)
	if err != nil {
		panic(err)
	}
	db = d
	return nil
}

func main() {
	// dataSourceName
	// dataSourceName := "host=localhost port=54321 user=postgres password=secret dbname=test sslmode=disable"
	dataSourceName := "postgres://postgres:secret@localhost:54321/test?sslmode=disable"

	err := Connect(dataSourceName)
	if err != nil {
		panic(err)
	}

	// insert
	// res, iner := db.Exec(`insert into tb (id, title, body, time) values($1, $2, $3, $4)`, 101, "test", "body content", 1234567891)

	// update
	// res, iner := db.Exec(`update tb set title=$1 where id=$2`, "test changed", 101)

	// select
	res, iner := db.Exec(`select * from tb`)
	if iner != nil {
		fmt.Println(iner)
		panic(iner)
	}

	lii, _ := res.LastInsertId()
	ra, _ := res.RowsAffected()
	fmt.Println(lii, ra)

	// err = db.Ping()
	// fmt.Println(err)

	fmt.Println("hello pg")
}
