package db

import (
	"context"
	"log"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5"
	"github.com/voyagesez/auservice/src/internals/db/sqlc"
)

type DatabaseInstance struct {
	RedisClient     *redis.Client
	PostgresqlQuery *sqlc.Queries
}

var once sync.Once
var databaseInstance *DatabaseInstance

func postgresqlQuery() (*sqlc.Queries, error) {
	ctx := context.Background()
	postgresqlDNS := "postgres://root:postgres@localhost:5432/voyagez_dev"
	conn, err := pgx.Connect(ctx, postgresqlDNS)
	if err != nil {
		return nil, err
	}
	log.Println("postgresql connected")
	return sqlc.New(conn), nil
}

func redisClient() (*redis.Client, error) {
	ctx := context.Background()
	redisDNS := "localhost:6379"
	client := redis.NewClient(&redis.Options{
		Addr:     redisDNS, // default
		Password: "",       // default
		DB:       0,        // default
	})
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	log.Println("redis connected")
	return client, nil
}

func createDatabaseInstance() (chan *redis.Client, chan *sqlc.Queries) {
	redisChan := make(chan *redis.Client)
	postgresqlChan := make(chan *sqlc.Queries)

	go func() {
		val, err := postgresqlQuery()
		if err != nil {
			log.Println("postgresql connection failed: ", err.Error())
			panic(err)
		}
		postgresqlChan <- val
	}()

	go func() {
		val, err := redisClient()
		if err != nil {
			log.Println("redis connection failed: ", err.Error())
			panic(err)
		}
		redisChan <- val
	}()
	return redisChan, postgresqlChan
}

func GetDatabaseInstance() *DatabaseInstance {
	return databaseInstance
}

func ConnectDatabase() {
	once.Do(func() {
		redisChan, postgresqlChan := createDatabaseInstance()
		databaseInstance = &DatabaseInstance{
			RedisClient:     <-redisChan,
			PostgresqlQuery: <-postgresqlChan,
		}
	})
}
