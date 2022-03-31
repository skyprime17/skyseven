package conf

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

var (
	dbUser = "root"
	dbPwd  = "123456"
	dbHost = "localhost"
	dbName = "skyseven"
)

func DbClient() (*sqlx.DB, error) {
	dataSource := getDataSource()
	client, err := sqlx.Connect("mysql", dataSource)
	if err != nil {
		fmt.Printf("error connecting")
		return nil, err
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxIdleConns(5)
	client.SetMaxOpenConns(10)
	return client, nil
}

func getDataSource() string {
	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbUser, dbPwd, dbHost, dbName)
}
