package mysql

import (
	"fmt"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lokmannicholas/weather-go/config"
)

type MYSQL struct {
	DB *sql.DB
}

var (
	db *sql.DB
)

func Get() *MYSQL {
	var err error
	if db == nil {
		if config.Get().MYSQL.Port == 0 {
			config.Get().MYSQL.Port = 3306
		}
		uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			config.Get().MYSQL.User,
			config.Get().MYSQL.Password,
			config.Get().MYSQL.Host,
			config.Get().MYSQL.Port,
			config.Get().MYSQL.DBName)

		db, err = sql.Open("mysql", uri)
		if err != nil {
			panic(err.Error())
		}
		db.SetMaxIdleConns(1)
		db.SetMaxOpenConns(10)
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	return &MYSQL{
		DB: db,
	}
}
