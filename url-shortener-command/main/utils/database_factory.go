package utils

import (
	"errors"
	"fmt"

	"github.com/go-pg/pg/v10"
)

const (
	DB_HOST       = "DB_HOST"
	DB_PORT       = "DB_PORT"
	CACHE_DB_HOST = "CACHE_DB_HOST"
	CACHE_DB_PORT = "CACHE_DB_PORT"
)

var (
	dbClient *pg.DB
)

type DatabaseFactory interface {
	CreateDb() (*pg.DB, error)
	GetDb() *pg.DB
	CloseDatabaseConnections()
}

type DatabaseFactoryImpl struct{}

func NewDatabaseFactory() *DatabaseFactoryImpl {
	return &DatabaseFactoryImpl{}
}

func (df *DatabaseFactoryImpl) CreateDb() (*pg.DB, error) {
	var err string

	dbClient = pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", "0.0.0.0", "5432"),
		Database: "db",
		User:     "postgres",
		Password: "postgres",
	})

	if len(err) > 0 {
		return nil, errors.New(err)
	}

	fmt.Println("DB client created successfully:", dbClient)
	return dbClient, nil
}

func (df *DatabaseFactoryImpl) GetDb() *pg.DB {
	return dbClient
}

func (df *DatabaseFactoryImpl) CloseDatabaseConnections() {
	if dbClient != nil {
		dbClient.Close()
	}
}
