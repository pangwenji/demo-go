package user

import (
	"context"
	v1 "demo/api/user/v1"
	"demo/service"
)

func (c *ControllerV1) IsSignedIn(ctx context.Context, req *v1.IsSignedInReq) (res *v1.IsSignedInRes, err error) {
	res = &v1.IsSignedInRes{
		OK: service.User().IsSignedIn(ctx),
	}
	return
}
