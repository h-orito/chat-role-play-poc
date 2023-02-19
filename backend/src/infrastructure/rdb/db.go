package db

import (
	"chat-role-play-poc/src/config"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DB struct {
	Connection *gorm.DB
}

func NewDB() DB {
	cfg, err := config.Load()
	if err != nil {
		panic(err.Error())
	}

	// https://github.com/go-sql-driver/mysql#examples
	connectStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true",
		cfg.Db.User,
		cfg.Db.Pass,
		cfg.Db.Host,
		cfg.Db.Name,
	)
	db, err := gorm.Open("mysql", connectStr)
	if err != nil {
		panic(err.Error())
	}

	return DB{
		Connection: db,
	}
}
