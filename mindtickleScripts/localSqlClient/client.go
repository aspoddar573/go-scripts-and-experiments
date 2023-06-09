package localSqlClient

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func GetSqlClient(schema string) *sql.DB {
	db, err := sql.Open("mysql",
		"root@tcp(127.0.0.1:3306)/"+schema)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
