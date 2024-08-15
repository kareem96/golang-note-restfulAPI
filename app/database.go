package app

import (
	"fmt"
	"time"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenConnectionDB(config Config) *gorm.DB  {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        config.DBUser,
        config.DBPassword,
        config.DBHost,
        config.DBPort,
        config.DBName,
    )
	// dialect := mysql.Open("root:developer@tcp(localhost:3306)/golang_note?charset=utf8mb4&parseTime=True&loc=Local")
	dialect := mysql.Open(dsn)
	db, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	})
	
	if err != nil{
		panic("failed to connect database")
	}

	sql, _ := db.DB()
	sql.SetMaxOpenConns(100)
	sql.SetMaxIdleConns(10)
	sql.SetConnMaxLifetime(30 * time.Minute)
	sql.SetConnMaxIdleTime(30 * time.Minute)

	return db
}