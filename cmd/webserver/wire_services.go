//go:build wireinject
// +build wireinject

package main

import (
	"context"
	authorRepository "go-echo-gorm-tempate/pkg/author/repository"
	authorRouter "go-echo-gorm-tempate/pkg/author/router"
	authorService "go-echo-gorm-tempate/pkg/author/service"
	bookRepository "go-echo-gorm-tempate/pkg/book/repository"
	bookRouter "go-echo-gorm-tempate/pkg/book/router"
	bookService "go-echo-gorm-tempate/pkg/book/service"
	"go-echo-gorm-tempate/pkg/config"

	"github.com/google/wire"
)

type services struct {
	authorRouter.AuthorRouterService
	bookRouter.BookRouterService
}

func initializeServices(ctx context.Context, conf *config.Config) (*services, error) {

	wire.Build(
		newDB,
		authorRepository.NewUserRepository,
		bookRepository.NewBookRepository,
		authorService.NewUserGettingService,
		bookService.NewBookGettingService,
		wire.Struct(new(authorRouter.AuthorRouterService), "*"),
		wire.Struct(new(bookRouter.BookRouterService), "*"),
		wire.Struct(new(services), "*"),
	)

	return nil, nil
}
