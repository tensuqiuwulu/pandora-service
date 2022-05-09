package mysql

import (
	"fmt"
	"strconv"
	"time"

	"github.com/tensuqiuwulu/pandora-service/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabaseConnection(configDatabase *config.Database) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configDatabase.Username,
		configDatabase.Password,
		configDatabase.Address,
		strconv.Itoa(int(configDatabase.Port)),
		configDatabase.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("Cannot connect to database: " + err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("Cannot connect to database: " + err.Error())
	}

	err = sqlDB.Ping()
	if err != nil {
		panic("Cannot ping the database: " + err.Error())
	}

	fmt.Println("Success connect to database")

	sqlDB.SetMaxIdleConns(int(configDatabase.MaxIdle))
	sqlDB.SetMaxOpenConns(int(configDatabase.MaxOpen))
	sqlDB.SetConnMaxLifetime(time.Minute * time.Duration(configDatabase.MaxLifeTime))
	sqlDB.SetConnMaxIdleTime(time.Minute * time.Duration(configDatabase.MaxIdleTime))
	return db
}

func Close(DB *gorm.DB) {
	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}

	err = sqlDB.Close()
	if err != nil {
		panic(err)
	}

	fmt.Println("database disconected")
}
