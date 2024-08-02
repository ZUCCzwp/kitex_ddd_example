// Package service -----------------------------
// @file      : user_service.go
// @author    : Jimmy
// @contact   : 2114272829@qq.com
// @time      : 2024/7/29 下午3:46
// -------------------------------------------
package service

import (
	"context"
	"github.com/ZUCCzwp/kitex_ddd_example/internal/domain/user/entity"
	"github.com/ZUCCzwp/kitex_ddd_example/internal/svc"
)

type UserServiceIer interface {
	Get(ctx context.Context, id int) (*entity.User, error)
}

type UserService struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserService(ctx context.Context, svcCtx *svc.ServiceContext) *UserService {
	return &UserService{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (u *UserService) Get(ctx context.Context, id int) (*entity.User, error) {
	data, err := u.svcCtx.UserRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
