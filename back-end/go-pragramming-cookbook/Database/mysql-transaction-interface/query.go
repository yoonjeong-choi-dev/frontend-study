package transaction

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func QueryTable(db DB, name string) error {
	query := fmt.Sprintf("SELECT name, created FROM %s where name=?", tableName)
	rows, err := db.Query(query, name)
	if err != nil {
		return err
	}
	defer func() { _ = rows.Close() }()

	fmt.Printf("Result with querying by name '%s'\n", name)
	for rows.Next() {
		var row DataRow
		if err := rows.Scan(&row.Name, &row.Created); err != nil {
			return err
		}

		fmt.Printf("\tName: %s\n\tCreated: %v\n", row.Name, row.Created)
	}

	return rows.Err()
}
