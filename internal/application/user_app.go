// Package application -----------------------------
// @file      : user_app.go
// @author    : Jimmy
// @contact   : 2114272829@qq.com
// @time      : 2024/7/29 下午3:46
// -------------------------------------------
package application

import (
	"context"
	"github.com/ZUCCzwp/kitex_ddd_example/internal/domain/entity"
	"github.com/ZUCCzwp/kitex_ddd_example/internal/svc"
)

type UserAppIer interface {
	Get(ctx context.Context, id int) (*entity.User, error)
}

type UserApp struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserApp(ctx context.Context, svcCtx *svc.ServiceContext) *UserApp {
	return &UserApp{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (u *UserApp) Get(ctx context.Context, id int) (*entity.User, error) {
	data, err := u.svcCtx.UserRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
