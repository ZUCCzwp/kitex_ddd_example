// Package repository -----------------------------
// @file      : user.go
// @author    : Jimmy
// @contact   : 2114272829@qq.com
// @time      : 2024/7/29 下午1:28
// -------------------------------------------
package repository

import (
	"context"
	"github.com/ZUCCzwp/kitex_ddd_example/internal/domain/user/entity"
)

//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOFILE

// UserRepository represent repository of the user
// Expect implementation by the infrastructure layer
type UserRepository interface {
	Get(ctx context.Context, id int) (*entity.User, error)
	GetAll(ctx context.Context) ([]*entity.User, error)
	Save(ctx context.Context, user *entity.User) error
}
