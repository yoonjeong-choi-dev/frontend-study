package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"mysql"
)

func main() {
	// .env 파일은 go.mod 와 동일한 위치(Goland 기준)
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	db, err := mysql.Setup()
	if err != nil {
		panic(err)
	}

	if err := mysql.ExecDatabaseExample(db); err != nil {
		panic(err)
	}

}
