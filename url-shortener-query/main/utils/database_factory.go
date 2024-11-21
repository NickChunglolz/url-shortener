package utils

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/redis/go-redis/v9"
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

type DatabaseFactory struct{
	config *Config
}

func NewDatabaseFactory(config *Config) *DatabaseFactory {
	return &DatabaseFactory{
		config: config,
	}
}

func (df *DatabaseFactory) CreateDb() (*pg.DB, error) {
	dbClient = pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", df.config.Downstream.Db.Host, df.config.Downstream.Db.Port),
		Database: df.config.Downstream.Db.Database,
		User:     df.config.Downstream.Db.User,
		Password: df.config.Downstream.Db.Password,
	})

	if err := dbClient.Ping(dbClient.Context()); err != nil {
		return nil, err
	}

	fmt.Println("Database connecion created successfully:", dbClient)
	return dbClient, nil
}

func (df *DatabaseFactory) CreateCacheDb() (*redis.Client, error) {
	cacheDbClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", df.config.Downstream.CacheDB.Host, df.config.Downstream.CacheDB.Port),
	})

	if err := cacheDbClient.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	fmt.Println("Cache database connecion created successfully:", cacheDbClient)
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
