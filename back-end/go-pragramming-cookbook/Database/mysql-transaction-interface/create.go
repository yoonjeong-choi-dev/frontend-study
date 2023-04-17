package transaction

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const tableName = "TestTable"

func CreateTable(db DB) error {
	query := fmt.Sprintf("CREATE TABLE %s (name VARCHAR(20), created DATETIME)", tableName)
	_, err := db.Exec(query)

	return err
}

func InsertData(db DB, name string) error {
	query := fmt.Sprintf(`INSERT INTO %s (name, created) VALUES ("%s", NOW())`, tableName, name)
	_, err := db.Exec(query)

	return err
}
