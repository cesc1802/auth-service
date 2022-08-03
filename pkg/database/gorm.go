package database

import (
	"fmt"
	"github.com/cesc1802/auth-service/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

type AppGormConfig struct {
	host     string
	port     string
	dbUser   string
	dbName   string
	password string
}

func NewAppGormConfig(host, port, dbUser, dbName, password string) *AppGormConfig {
	return &AppGormConfig{
		host:     host,
		port:     port,
		dbUser:   dbUser,
		dbName:   dbName,
		password: password,
	}
}

func (cnf *AppGormConfig) Uri() string {
	//user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cnf.dbUser, cnf.password, cnf.host, cnf.port, cnf.dbName)
}

type appGorm struct {
	db     *gorm.DB
	logger *logger.Logger
}

func NewAppGorm(cnf *AppGormConfig) *appGorm {
	dbGormLogger := DefaultLogger("gorm")
	dbGormLogger.LogLevel = gormLogger.Info

	db, err := gorm.Open(mysql.Open(cnf.Uri()), &gorm.Config{
		Logger:                                   dbGormLogger,
		PrepareStmt:                              true,
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
	})

	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetConnMaxLifetime(1 * time.Hour)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	if err != nil {
		return nil
	}
	return &appGorm{
		db:     db,
		logger: dbGormLogger.ZapLogger,
	}
}

func (db *appGorm) GetDB() *gorm.DB {
	if db != nil {
		return db.db
	}
	return nil
}
