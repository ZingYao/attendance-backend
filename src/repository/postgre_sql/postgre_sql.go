package postgre_sql

import (
	"attendance_backend/src/dbsource"
	"attendance_backend/src/repository"
	"gorm.io/gorm"
)

type datasource struct {
	dbConn gorm.DB
}

func GetRepository() repository.Factory {
	dbConn := dbsource.GetPostgreSqlConn()
	return &datasource{dbConn: dbConn}
}
