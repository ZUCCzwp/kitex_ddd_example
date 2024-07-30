package interfaces

import (
	"context"
	"github.com/ZUCCzwp/kitex_ddd_example/internal/domain/service"
	"github.com/ZUCCzwp/kitex_ddd_example/internal/svc"
	user "github.com/ZUCCzwp/kitex_ddd_example/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	SvcCtx *svc.ServiceContext
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *user.Request) (resp *user.Response, err error) {
	l := service.NewUserService(ctx, s.SvcCtx)
	data, err := l.Get(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	return &user.Response{Name: data.Name}, nil
}
