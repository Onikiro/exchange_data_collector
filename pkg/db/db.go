package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func AddConfig(symbol string) bool {
	db, err := sql.Open("sqlite3", "database.db")
	logErr(err)
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO configs(symbol) values(?)")
	logErr(err)

	res, err := stmt.Exec(symbol)
	logErr(err)

	count, err := res.RowsAffected()
	logErr(err)

	return count > 0
}

func GetConfigs() []string {
	db, err := sql.Open("sqlite3", "database.db")
	logErr(err)
	defer db.Close()

	var count int
	countRows, err := db.Query("SELECT COUNT(*) FROM configs")
	logErr(err)
	if countRows.Next() {
		err = countRows.Scan(&count)
		logErr(err)
	}

	countRows.Close()

	rows, err := db.Query("SELECT * FROM configs")
	logErr(err)
	var symbol string
	var symbols []string = make([]string, count)
	i := 0
	for rows.Next() {
		err = rows.Scan(&symbol)
		logErr(err)
		symbols[i] = symbol
		i++
	}

	rows.Close()
	return symbols
}

func DeleteConfig(symbol string) bool {
	db, err := sql.Open("sqlite3", "database.db")
	logErr(err)
	defer db.Close()

	stmt, err := db.Prepare("delete from configs where symbol=?")
	logErr(err)

	res, err := stmt.Exec(symbol)
	logErr(err)

	count, err := res.RowsAffected()
	logErr(err)

	return count > 0
}

func logErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
