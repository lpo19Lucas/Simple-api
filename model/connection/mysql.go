package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var connectionDb *sql.DB

// MySqlConnect return connection with db
func MySqlConnect() *sql.DB {
	if connectionDb == nil {
		fmt.Println("Connect in DB")
		connectionDb, _ = sql.Open("mysql", "root:123@/pfm")
		fmt.Println(connectionDb.Ping())
		if connectionDb.Ping() != nil {
			fmt.Println(connectionDb.Ping())
		}
	}
	return connectionDb
}
