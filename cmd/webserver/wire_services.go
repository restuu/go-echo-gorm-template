//go:build wireinject
// +build wireinject

package main

import (
	"context"
	authorRepository "go-echo-gorm-tempate/pkg/author/repository"
	authorRouter "go-echo-gorm-tempate/pkg/author/router"
	authorService "go-echo-gorm-tempate/pkg/author/service"
	"go-echo-gorm-tempate/pkg/config"

	"github.com/google/wire"
)

type services struct {
	authorRouter.AuthorRouterService
}

func initializeServices(ctx context.Context, conf *config.Config) (*services, error) {

	wire.Build(
		newDB,
		authorRepository.NewUserRepository,
		authorService.NewUserGettingService,
		wire.Struct(new(authorRouter.AuthorRouterService), "*"),
		wire.Struct(new(services), "*"),
	)

	return nil, nil
}
