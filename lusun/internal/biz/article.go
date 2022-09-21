package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"lusun/internal/data/po"
)

// ArticleRepo is a Greater repo.
type ArticleRepo interface {
	Detail(id int64, article *po.ZbpPost) error
}

type ArticleUseCase struct {
	repo ArticleRepo
	log  *log.Helper
}

func NewArticleUseCase(repo ArticleRepo) *ArticleUseCase {
	return &ArticleUseCase{
		repo: repo,
	}
}

func (a *ArticleUseCase) Detail(id int64) (*po.ZbpPost, error) {
	article := &po.ZbpPost{}
	err := a.repo.Detail(id, article)
	if err != nil {
		return nil, err
	}

	return article, err
}
