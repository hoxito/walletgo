package db

import (
	"fmt"
	"log"
	"time"

	"database/sql"

	"github.com/go-redis/redis/v7"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hoxito/walletgo/tools/env"
)

func Mysql() *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s?columnsWithAlias=true&parseTime=true", env.Get().MysqlURL))
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	fmt.Println("Connected to MysqlDB: ")
	return db
}

func Redis() *redis.Client {
	rd := redis.NewClient(&redis.Options{
		Addr:     env.Get().RedisURL,
		Password: "",
		DB:       0,
	})
	return rd

}

//Client instance
var DB *sql.DB = Mysql()
var RedisDB *redis.Client = Redis()

//TODO func IsUniqueKeyError(err error) bool
