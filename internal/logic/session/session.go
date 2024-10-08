package session

import (
	"context"
	"demo/internal/consts"
	"demo/model/entity"
	"demo/service"
)

type (
	sSession struct{}
)

func init() {
	service.RegisterSession(New())
}

func New() service.ISession {
	return &sSession{}
}

// SetUser sets user into the session.
func (s *sSession) SetUser(ctx context.Context, user *entity.User) error {
	return service.BizCtx().Get(ctx).Session.Set(consts.UserSessionKey, user)
}

// GetUser retrieves and returns the user from session.
// It returns nil if the user did not sign in.
func (s *sSession) GetUser(ctx context.Context) *entity.User {
	customCtx := service.BizCtx().Get(ctx)
	if customCtx != nil {
		if v := customCtx.Session.MustGet(consts.UserSessionKey); !v.IsNil() {
			var user *entity.User
			_ = v.Struct(&user)
			return user
		}
	}
	return nil
}

// RemoveUser removes user rom session.
func (s *sSession) RemoveUser(ctx context.Context) error {
	customCtx := service.BizCtx().Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Remove(consts.UserSessionKey)
	}
	return nil
}
