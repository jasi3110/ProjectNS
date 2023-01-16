package utls

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DbConnection *sql.DB

func CreateDbConnection() (*sql.DB, bool) {

	err := DbConnection.Ping()
	if err != nil {
		fmt.Println("DB Connection Ping Failed", err)
	}

	_, isOpen := OpenDbConnection()

	if isOpen {
		return DbConnection, true
	} else {
		return DbConnection, false
	}
}

func OpenDbConnection() (*sql.DB, bool) {
	// const (
	// 	host     = "localhost"
	// 	port     = 5432
	// 	user     = "postgres"
	// 	password = "root"
	// 	dbname   = "postgres"
	// )
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)

	myDb, err := sql.Open("postgres", Config.ConnectionString)

	if err != nil {
		fmt.Println("Db Open Failed:", err)
	}

	myDb.SetMaxOpenConns(0)
	myDb.SetMaxIdleConns(100)
	DbConnection = myDb
	return myDb, true
}

func init() {
	log.Println("Connecting to postgrs db...")
	OpenDbConnection()
}
