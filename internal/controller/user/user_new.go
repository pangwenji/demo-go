package user

import (
	"demo/api/user"
)

type ControllerV1 struct{}

func NewV1() user.IUserV1 {
	return &ControllerV1{}
}
