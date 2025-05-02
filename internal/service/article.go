package service

import (
	"context"
	"github.com/LEILEI0628/GoWeb/internal/domain"
	"github.com/LEILEI0628/GoWeb/internal/repository/article"
)

type ArticleServiceInterface interface {
	Save(ctx context.Context, article domain.Article) (int64, error)
	Publish(ctx context.Context, art domain.Article) (int64, error)
}
type ArticleService struct {
	repo article.ArticleRepository
}

func NewArticleService(repo article.ArticleRepository) ArticleServiceInterface {
	return &ArticleService{repo: repo}
}

func (svc *ArticleService) Save(ctx context.Context, article domain.Article) (int64, error) {
	if article.Id > 0 {
		// 修改帖子（有Id）
		err := svc.repo.Update(ctx, article)
		return article.Id, err
	}
	return svc.repo.Create(ctx, article) // 新建帖子
}

func (svc *ArticleService) Publish(ctx context.Context, art domain.Article) (int64, error) {
	// 制作库
	//id, err := a.repo.Create(ctx, art)
	//// 线上库
	//a.repo.SyncToLiveDB(ctx, art)
	panic("implement me")
}
