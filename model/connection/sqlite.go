package connection

import (
	"database/sql"
	"fmt"
)

const PATH_FILE_DB = "./db/database.db"

var connectionDb *sql.DB

// SqliteConnect return connection with db
func SqliteConnect() *sql.DB {
	if connectionDb == nil {
		fmt.Println("Connect in DB")
		connectionDb, _ = sql.Open("sqlite3", PATH_FILE_DB)
		// if error != nil {
		// 	fmt.Println("ERRO", error)
		// }
	}
	return connectionDb
}
