package service

import (
	"context"
	"github.com/LEILEI0628/GoWeb/internal/domain"
	"github.com/LEILEI0628/GoWeb/internal/repository"
)

type ArticleServiceInterface interface {
	Save(ctx context.Context, article domain.Article) (int64, error)
}
type ArticleService struct {
	repo repository.ArticleRepository
}

func NewArticleService(repo repository.ArticleRepository) *ArticleService {
	return &ArticleService{repo: repo}
}

func (s *ArticleService) Save(ctx context.Context, article domain.Article) (int64, error) {
	if article.Id > 0 {
		// 修改帖子（有Id）
		err := s.repo.Update(ctx, article)
		return article.Id, err
	}
	return s.repo.Create(ctx, article) // 新建帖子
}
