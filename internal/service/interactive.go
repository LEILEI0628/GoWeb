package service

import (
	"context"
	loggerx "github.com/LEILEI0628/GinPro/middleware/logger"
	"github.com/LEILEI0628/GoWeb/internal/domain"
	"github.com/LEILEI0628/GoWeb/internal/repository"
	"golang.org/x/sync/errgroup"
)

type InteractiveServiceInterface interface {
	IncrReadCnt(ctx context.Context, biz string, bizId int64) error
	// Like 点赞
	Like(ctx context.Context, biz string, bizId int64, uid int64) error
	// CancelLike 取消点赞
	CancelLike(ctx context.Context, biz string, bizId int64, uid int64) error
	// Collect 收藏, cid 是收藏夹的 ID
	// cid 不一定有，或者说 0 对应的是该用户的默认收藏夹
	Collect(ctx context.Context, biz string, bizId, cid, uid int64) error
	Get(ctx context.Context, biz string, bizId, uid int64) (domain.Interactive, error)
}

type InteractiveService struct {
	repo repository.InteractiveRepositoryInterface
	l    loggerx.Logger
}

func NewInteractiveService(repo repository.InteractiveRepositoryInterface,
	l loggerx.Logger) InteractiveServiceInterface {
	return &InteractiveService{
		repo: repo,
		l:    l,
	}
}

func (svc *InteractiveService) IncrReadCnt(ctx context.Context, biz string, bizId int64) error {
	return svc.repo.IncrReadCnt(ctx, biz, bizId)
}

func (svc *InteractiveService) Get(ctx context.Context,
	biz string, bizId, uid int64) (domain.Interactive, error) {
	// 按照 repository 的语义(完成 domain.Interactive 的完整构造)，你这里拿到的就应该是包含全部字段的
	var (
		eg        errgroup.Group
		intr      domain.Interactive
		liked     bool
		collected bool
	)
	eg.Go(func() error {
		var err error
		intr, err = svc.repo.Get(ctx, biz, bizId)
		return err
	})
	eg.Go(func() error {
		var err error
		liked, err = svc.repo.Liked(ctx, biz, bizId, uid)
		return err
	})
	eg.Go(func() error {
		var err error
		liked, err = svc.repo.Collected(ctx, biz, bizId, uid)
		return err
	})
	err := eg.Wait()
	if err != nil {
		return domain.Interactive{}, err
	}
	intr.Liked = liked
	intr.Collected = collected
	return intr, err
}

func (svc *InteractiveService) Like(ctx context.Context, biz string, bizId int64, uid int64) error {
	// 点赞
	return svc.repo.IncrLike(ctx, biz, bizId, uid)
}

func (svc *InteractiveService) CancelLike(ctx context.Context, biz string, bizId int64, uid int64) error {
	return svc.repo.DecrLike(ctx, biz, bizId, uid)
}

// Collect 收藏
func (svc *InteractiveService) Collect(ctx context.Context,
	biz string, bizId, cid, uid int64) error {
	// service 还叫做收藏
	// repository
	return svc.repo.AddCollectionItem(ctx, biz, bizId, cid, uid)
}
