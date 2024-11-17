package utils

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/redis/go-redis/v9"
)

const (
	DB_HOST       = "DB_HOST"
	DB_PORT       = "DB_PORT"
	CACHE_DB_HOST = "CACHE_DB_HOST"
	CACHE_DB_PORT = "CACHE_DB_PORT"
)

var (
	dbClient      *pg.DB
	cacheDbClient *redis.Client
)

type DatabaseFactoryInterface interface {
	CreateDb() (*pg.DB, error)
	CreateCacheDb() (*redis.Client, error)
	GetDb() *pg.DB
	GetCacheDb() *redis.Client
	CloseDatabaseConnections()
}

type DatabaseFactory struct{}

func NewDatabaseFactory() *DatabaseFactory {
	return &DatabaseFactory{}
}

func (df *DatabaseFactory) CreateDb() (*pg.DB, error) {
	dbClient = pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", "0.0.0.0", "5432"),
		Database: "db",
		User:     "postgres",
		Password: "postgres",
	})

	if err := dbClient.Ping(dbClient.Context()); err != nil {
		return nil, err
	}

	fmt.Println("Database connecion created successfully:", dbClient)
	return dbClient, nil
}

func (df *DatabaseFactory) CreateCacheDb() (*redis.Client, error) {
	cacheDbClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", "0.0.0.0", "6379"),
	})

	if err := cacheDbClient.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	fmt.Println("Cache database connecion created successfully:", dbClient)
	return cacheDbClient, nil
}

func (df *DatabaseFactory) GetDb() *pg.DB {
	return dbClient
}

func (df *DatabaseFactory) GetCacheDb() *redis.Client {
	return cacheDbClient
}

func (df *DatabaseFactory) CloseDatabaseConnections() {
	if dbClient != nil {
		dbClient.Close()
	}

	if cacheDbClient != nil {
		cacheDbClient.Close()
	}
}
