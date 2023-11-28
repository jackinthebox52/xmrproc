package db

import (
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

type Transaction struct {
	UUID          string
	Amount        float64
	FromAddress   string
	WalletAddress string
	TimeCreated   string
	TimeConfirmed string
}

func initializeDatabase() {
	//create new sqlite database if it doesn't exist
	var err error
	DB, err = sql.Open("sqlite3", fmt.Sprintf("%s/transactions.db", APP_DATA_DIR))
	if err != nil {
		log.Fatal(err)
	}
}
