package datastore

import (
	"context"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DatabaseURL string type alias.
type DatabaseURL string

func Open(ctx context.Context, url DatabaseURL) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		// DSN:             "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // data source name
		DSN:                       string(url),
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index,
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})

	return db, err
}
