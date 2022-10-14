package dbsource

import (
	"attendance_backend/src/entities"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreSql struct {
}

var db *gorm.DB

func (PostgreSql) InitPostgreSql(pgConf entities.PostgreSqlStruct) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		pgConf.Host, pgConf.User, pgConf.Pass, pgConf.DB, pgConf.Port, pgConf.TimeZone)
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("connect to postgreSql failed,%v", err))
	}
}

func GetPostgreSqlConn() gorm.DB {
	return *db
}
