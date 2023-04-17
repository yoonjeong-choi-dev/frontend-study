package transaction

import "fmt"

func ExecDatabaseExample(db DB) error {
	defer func() { _, _ = db.Exec(fmt.Sprintf("DROP TABLE %s", tableName)) }()

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
