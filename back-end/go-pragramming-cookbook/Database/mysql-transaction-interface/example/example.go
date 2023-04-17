package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"transaction"
)

func main() {
	// .env 파일은 go.mod 와 동일한 위치(Goland 기준)
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db, err := transaction.Setup()
	if err != nil {
		panic(err)
	}

	// init transaction
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	// If the process stopped before commit, rollback
	defer func() { _ = tx.Rollback() }()

	if err := transaction.ExecDatabaseExample(tx); err != nil {
		panic(err)
	}

	// commit
	if err := tx.Commit(); err != nil {
		panic(err)
	}
}
