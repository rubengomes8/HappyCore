package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" //no lint
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

func NewConfig(
	host,
	database,
	user,
	password string,
	port int) Config {
	return Config{
		Host:     host,
		Port:     port,
		Database: database,
		User:     user,
		Password: password,
	}
}

func InitDB(c Config) (*sql.DB, error) {

	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.Database,
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitGormDB(c Config) (*gorm.DB, error) {

	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.Database,
	)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
