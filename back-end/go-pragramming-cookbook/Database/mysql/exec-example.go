package mysql

import (
	"database/sql"
	"fmt"
)

func ExecDatabaseExample(db *sql.DB) error {
	defer db.Exec(fmt.Sprintf("DROP TABLE %s", TableName))

	if err := CreateTable(db); err != nil {
		return err
	}

	if err := InsertData(db, "Yoonjeong"); err != nil {
		return err
	}

	if err := InsertData(db, "Yoonjeong"); err != nil {
		return err
	}

	if err := QueryTable(db, "Yoonjeong"); err != nil {
		return err
	}

	if err := QueryTable(db, "NoData"); err != nil {
		return err
	}

	return nil
}
