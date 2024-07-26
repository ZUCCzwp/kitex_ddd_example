package handler

import (
	"context"
	"github.com/ZUCCzwp/ddd/my-awesome-service/internal/svc"
	"github.com/ZUCCzwp/ddd/my-awesome-service/kitex_gen/hello"
	"github.com/ZUCCzwp/ddd/my-awesome-service/pkg/onelog"
)

// HelloServiceImpl implements the last service interface defined in the IDL.
type HelloServiceImpl struct {
	SvcCtx *svc.ServiceContext
}

func (h HelloServiceImpl) SayHello(ctx context.Context, req *hello.Request) (res *hello.Response, err error) {
	res = &hello.Response{Message: req.Message}
	onelog.CtxDebugf(ctx, "SayHello req:%v", req)
	return
}
