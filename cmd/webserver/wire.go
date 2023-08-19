//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"go-echo-gorm-tempate/adapter/config"
	"go-echo-gorm-tempate/app/author"
	"go-echo-gorm-tempate/app/book"

	"github.com/google/wire"
)

var (
	usecasesProvider = wire.NewSet(
		author.NewAuthorUsecase,
		book.NewBookUsecase,
		wire.Struct(new(Usecases), "*"),
	)
)

func Init(ctx context.Context) (*App, error) {
	wire.Build(
		config.NewConfig,
		newDB,
		author.NewAuthorRepository,
		book.NewBookRepository,
		usecasesProvider,
		wire.Struct(new(App), "Usecases", "Config", "Context"),
	)

	return &App{}, nil
}
