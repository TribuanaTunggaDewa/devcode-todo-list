package database

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Db interface {
	Init(ctx context.Context) (any, error)
	GetProvider() string
}

type db struct {
	Host        string
	User        string
	Pass        string
	Port        string
	Name        string
	Provider    string
	MaxOpenConn int
	MaxIdleConn int
	MaxLifetime int
}

type dbMysql struct {
	db
	Charset     string
	ParseTime   string
	Loc         string
	AutoMigrate bool
}

func (c *dbMysql) Init(ctx context.Context) (any, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s", c.User, c.Pass, c.Host, c.Port, c.Name, c.Charset, c.ParseTime, c.Loc)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetConnMaxLifetime(time.Duration(c.MaxLifetime) * time.Minute)
	sqlDB.SetMaxOpenConns(c.MaxOpenConn)
	sqlDB.SetMaxIdleConns(c.MaxIdleConn)

	if c.AutoMigrate {
		logrus.Info("auto migrate ...")
		err = db.AutoMigrate(Model...)
		if err != nil {
			db.DisableForeignKeyConstraintWhenMigrating = true
			time.Sleep(5 * time.Second)
			db.AutoMigrate(Model...)
			db.DisableForeignKeyConstraintWhenMigrating = false
			time.Sleep(5 * time.Second)
			db.AutoMigrate(Model...)
		}
	}

	return db, nil
}

func (c *dbMysql) GetProvider() string {
	return c.Provider
}
