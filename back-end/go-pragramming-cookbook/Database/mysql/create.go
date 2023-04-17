package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const TableName = "TestTable"

func CreateTable(db *sql.DB) error {
	query := fmt.Sprintf("CREATE TABLE %s (name VARCHAR(20), created DATETIME)", TableName)
	_, err := db.Exec(query)

	return err
}

func InsertData(db *sql.DB, name string) error {
	query := fmt.Sprintf(`INSERT INTO %s (name, created) VALUES ("%s", NOW())`, TableName, name)
	_, err := db.Exec(query)

	return err
}
