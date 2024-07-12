package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	DBHost string `yaml:"dbHost"`
	DBPort string `yaml:"dbPort"`
	DBUser string `yaml:"dbUser"`
	DBPass string `yaml:"dbPass"`
	Dbase  string `yaml:"dbase"`
}

func LoginDB(db *DBConfig) (*sql.DB, error) {
	dbConn := fmt.Sprintf("%s:%s@tcp(%s)/%s", db.DBUser, db.DBPass, db.DBHost, db.Dbase)
	database, err := sql.Open("mysql", dbConn)
	if err != nil {
		return nil, err
	}
	return database, nil
}
