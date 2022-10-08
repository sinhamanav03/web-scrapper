package db

import (
	"database/sql"
	"fmt"
)

type Config struct {
	Host     string
	Name     string
	User     string
	Password string
	Port     int
}

func NewPostgresDB(conf *Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", conf.User, conf.Password, conf.Host, conf.Port, conf.Name)

	connDB, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to open db connection: %w", err)
	}

	err = connDB.Ping()

	if err != nil {
		return nil, fmt.Errorf("unable to ping to db: %w", err)
	}

	return connDB, nil
}
