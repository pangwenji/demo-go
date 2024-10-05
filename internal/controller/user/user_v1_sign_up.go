package user

import (
	"context"
	v1 "demo/api/user/v1"
	"demo/model"
	"demo/service"
)

func (c *ControllerV1) SignUp(ctx context.Context, req *v1.SignUpReq) (res *v1.SignUpRes, err error) {
	err = service.User().Create(ctx, model.UserCreateInput{
		Passpost: req.Passport,
		Password: req.Password,
		Nickname: req.Nickname,
	})
	return
}
