package database

import (
	"context"
	"errors"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var dbConnectionsMysql *gorm.DB

func Init(ctx context.Context) {
	dbConnectionsMysql = &gorm.DB{}
	mysqlDb := dbMysql{
		db: db{
			Host:        os.Getenv("MYSQL_HOST"),
			Name:        os.Getenv("MYSQL_DBNAME"),
			Port:        os.Getenv("MYSQL_PORT"),
			User:        os.Getenv("MYSQL_USER"),
			Pass:        os.Getenv("MYSQL_PASSWORD"),
			Provider:    "mysql",
			MaxOpenConn: 200,
			MaxIdleConn: 10,
			MaxLifetime: 15,
		},
		AutoMigrate: true,
		Charset:     "utf8mb4",
		ParseTime:   "True",
		Loc:         "Local",
	}
	db, err := mysqlDb.Init(ctx)
	if err != nil {
		logrus.Info(err)
		panic("failed to connect to database")
	}

	dbConnectionsMysql = db.(*gorm.DB)
}

func getConnection() (any, error) {
	if dbConnectionsMysql == nil {
		return nil, errors.New("connection is undefined")
	}
	return dbConnectionsMysql, nil
}
