package main

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/mustafasegf/go-shortener/api"
	"github.com/mustafasegf/go-shortener/entity"
	"github.com/mustafasegf/go-shortener/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// err := util.SetLogger()
	// if err != nil {
	// 	log.Fatal("cannot set logger: ", err)
	// }

	config, err := util.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: config.DBURL,
	}))
	if err != nil {
		log.Fatal("canot load db: ", err)
	}

	// ping db
	dbsql, err := db.DB()
	if err != nil {
		log.Fatal("cannot get db instance")
	}

	err = dbsql.Ping()
	if err != nil {
		log.Fatal("cannot ping db: ", err)
	}

	err = db.AutoMigrate(&entity.LinkModel{})
	if err != nil {
		log.Fatal("cannot migrate db: ", err)
	}

	opt, err := redis.ParseURL(config.RedisURL)
	if err != nil {
		log.Fatal("cannot parse redis url: ", err)
	}

	rdb := redis.NewClient(opt)

	if err = rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatal("cannot ping redis: ", err)
	}

	server := api.MakeServer(config, db, rdb)
	server.RunServer()
}
