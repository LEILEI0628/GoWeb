package dao

import (
	"context"
	"github.com/LEILEI0628/GoWeb/internal/repository/dao/po"
	"gorm.io/gorm"
	"time"
)

type ArticleDAO interface {
	Insert(ctx context.Context, article po.Article) (int64, error)
	UpdateById(ctx context.Context, article po.Article) error
}

func NewArticleDAO(db *gorm.DB) ArticleDAO {
	return &GORMArticleDAO{
		db: db,
	}
}

type GORMArticleDAO struct {
	db *gorm.DB
}

func (dao *GORMArticleDAO) Insert(ctx context.Context, article po.Article) (int64, error) {
	now := time.Now().UnixMilli()
	article.CreateTime = now
	article.UpdateTime = now
	err := dao.db.WithContext(ctx).Create(&article).Error
	return article.Id, err
}

func (dao *GORMArticleDAO) UpdateById(ctx context.Context, article po.Article) error {
	now := time.Now().UnixMilli()
	article.UpdateTime = now
	// 依赖GORM忽略零值的特性，默认会用主键进行更新（不推荐，可读性很差）
	//err := dao.db.WithContext(ctx).Updates(&article).Error
	err := dao.db.WithContext(ctx).Model(&article).
		Where("id=?", article.Id).
		Updates(map[string]any{
			"title":       article.Title,
			"content":     article.Content,
			"update_time": article.UpdateTime,
		}).Error
	return err
}
