package v1

import (
	"demo/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type ProfileReq struct {
	g.Meta `path:"/user/profile" method:"get" tags:"UserService" summary:"Get the profile of current user"`
}
type ProfileRes struct {
	*entity.User
}
