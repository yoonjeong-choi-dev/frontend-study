package connectionpools

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func Setup() (*sql.DB, error) {
	host := fmt.Sprintf("%s:%s@/%s?parseTime=true",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DB"),
	)
	db, err := sql.Open("mysql", host)

	if err != nil {
		return nil, err
	}

	fmt.Printf("Success to connect to %s\n", host)

	// max connections: 24
	db.SetMaxOpenConns(24)
	db.SetMaxIdleConns(24)

	return db, nil
}
