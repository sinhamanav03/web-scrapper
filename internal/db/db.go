package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Name     string
	User     string
	Password string
	Port     int
}

func NewPostgresDB(conf *Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable ", conf.Host, conf.User, conf.Password, conf.Name, conf.Port)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}

	return db, nil

}
