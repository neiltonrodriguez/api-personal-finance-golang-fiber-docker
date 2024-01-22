package database

import (
	"database/sql"
	"fmt"
	"personal-finance-api/domain"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB() (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", domain.GlobalConfig.Username, domain.GlobalConfig.Password, domain.GlobalConfig.Host, domain.GlobalConfig.Port, domain.GlobalConfig.Database)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
