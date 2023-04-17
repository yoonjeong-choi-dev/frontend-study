package transaction

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

type DataRow struct {
	Name    string
	Created *time.Time
}

func Setup() (*sql.DB, error) {
	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@/%s?parseTime=true",
			os.Getenv("MYSQL_USER"),
			os.Getenv("MYSQL_PASSWORD"),
			os.Getenv("MYSQL_DB")),
	)

	if err != nil {
		return nil, err
	}

	return db, err
}
