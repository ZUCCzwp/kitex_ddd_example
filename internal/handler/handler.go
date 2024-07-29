package handler

import (
	"context"
	"github.com/ZUCCzwp/kitex_ddd_example/internal/svc"
	"github.com/ZUCCzwp/kitex_ddd_example/kitex_gen/hello"
	"github.com/ZUCCzwp/kitex_ddd_example/pkg/onelog"
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
