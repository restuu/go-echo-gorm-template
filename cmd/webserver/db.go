package main

import (
	"context"
	"go-echo-gorm-tempate/adapter/config"
	"go-echo-gorm-tempate/adapter/datastore"

	"gorm.io/gorm"
)

func newDB(ctx context.Context, conf *config.Config) (db *gorm.DB, err error) {
	return datastore.Open(ctx, datastore.DatabaseURL(conf.DatabaseURL))
}
