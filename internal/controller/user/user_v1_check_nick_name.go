package user

import (
	"context"
	v1 "demo/api/user/v1"
	"demo/service"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c ControllerV1) CheckNickName(ctx context.Context, req *v1.CheckNickNameReq) (res *v1.CheckNickNameRes, err error) {
	available, err := service.User().IsNicknameAvailable(ctx, req.Nickname)
	if err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.Newf(`Nickname "%s" is already token by others`, req.Nickname)
	}
	return
}
