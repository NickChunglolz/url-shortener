package utils

import (
	"errors"
	"fmt"

	"github.com/go-pg/pg/v10"
)

var (
	dbClient *pg.DB
)

type DatabaseFactory interface {
	CreateDb() (*pg.DB, error)
	GetDb() *pg.DB
	CloseDatabaseConnections()
}

type DatabaseFactoryImpl struct{
	config *Config
}

func NewDatabaseFactory(config *Config) *DatabaseFactoryImpl {
	return &DatabaseFactoryImpl{
		config: config,
	}
}

func (df *DatabaseFactoryImpl) CreateDb() (*pg.DB, error) {
	var err string

	dbClient = pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", df.config.Downstream.Db.Host, df.config.Downstream.Db.Port),
		Database: df.config.Downstream.Db.Database,
		User:     df.config.Downstream.Db.User,
		Password: df.config.Downstream.Db.Password,
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
