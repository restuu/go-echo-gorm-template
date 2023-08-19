package author

import (
	"context"
	"go-echo-gorm-tempate/domain"
)

// FindAll implements AuthorService.
func (svc *Usecase) FindAll(ctx context.Context) ([]domain.Author, error) {
	return svc.authorRepository.FindAll(ctx)
}
