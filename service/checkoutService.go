package service

import (
	"context"
	"eniqilo_store/model/web"
)

type CheckoutService interface {
	Create(ctx context.Context, req web.CheckoutReq) error
}
