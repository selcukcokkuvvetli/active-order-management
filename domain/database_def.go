package domain

import "database/sql"

type Database interface {
	Open() (*sql.DB, error)
	Close(db *sql.DB)
	Migrate(db *sql.DB) error
}

type DatabaseContext struct {
	Host     string
	Name     string
	Port     int
	Password string
	UserName string
}

func NewDatabaseContext(host, name, userName, password string, port int) Database {
	return &DatabaseContext{
		Host:     host,
		Name:     name,
		Port:     port,
		Password: password,
		UserName: userName,
	}
}