package database

import (
	"database/sql"
	"fmt"
)

type ConnectionSettings struct {
	Host     string
	Port     int
	Username string
	DBName   string
	SSLMode  string
	Password string
}

func NewPosgresqlConnection(settings ConnectionSettings) (*sql.DB, error) {
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
			settings.Host, settings.Port, settings.Username, settings.DBName, settings.SSLMode, settings.Password),
	)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
