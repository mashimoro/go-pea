package database

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func (d Database) GetConnectionDB() *gorm.DB {
	if d.DB == nil {
		d.DB = d.initConnect()
	}
	return d.DB
}

func (d Database) initConnect() *gorm.DB {
	var err error
	dsn := "sqlserver://sa:yourStrong(!)Password@localhost:1433?database=master"

	d.DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("database connection failed : " + err.Error())
	} else {
		fmt.Println("database connection success")
	}

	// d.DB.AutoMigrate(&purchasing_status.PurchasingStatus{}, user.User{})

	return d.DB
}
