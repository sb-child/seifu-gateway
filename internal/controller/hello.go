package controller

import (
	"context"

	v1 "seifu-gw/api/v1"
	// "github.com/gogf/gf/v2/frame/g"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (c *cHello) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	// g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	res = &v1.HelloRes{}
	return
}
