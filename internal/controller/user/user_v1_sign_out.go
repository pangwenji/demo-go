package user

import (
	"context"
	v1 "demo/api/user/v1"
	"demo/service"
)

func (c *ControllerV1) SignOut(ctx context.Context, req *v1.SignOutReq) (res *v1.SignOutRes, err error) {
	err = service.User().SignOut(ctx)
	return
}
