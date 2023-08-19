package author

import (
	"go-echo-gorm-tempate/domain"
	"sync"
)

var (
	authorService     domain.AuthorUsecase
	authorServiceOnce sync.Once
)

type Usecase struct {
	authorRepository domain.AuthorRepository
}

func NewAuthorUsecase(authorRepository domain.AuthorRepository) domain.AuthorUsecase {
	authorServiceOnce.Do(func() {
		authorService = &Usecase{
			authorRepository: authorRepository,
		}
	})

	return authorService
}
