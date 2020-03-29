package core

import (
	"database/sql"
	"fmt"
	rd "github.com/go-redis/redis/v7"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var redis *rd.Client

func Init() {

	db, _ = sql.Open("mysql", "root:system@tcp(127.0.0.1:3306)/employees?charset=utf8&parseTime=true")
	err := db.Ping()
	fmt.Println(db, err)

	redis = rd.NewClient(&rd.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := redis.Ping().Result()
	fmt.Println(pong, err)
}