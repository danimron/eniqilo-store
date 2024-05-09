package service

import (
	"context"
	"eniqilo_store/model/web"
)

type StaffService interface {
	Register(ctx context.Context, request web.StaffRegisterReq) (web.StaffRes, error)
	Login(ctx context.Context, request web.StaffLoginReq) (web.StaffRes, error)
}
