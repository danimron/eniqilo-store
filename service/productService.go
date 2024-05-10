package service

import (
	"context"
	"eniqilo_store/model/web"
)

type ProductService interface {
	Create(ctx context.Context, request web.ProductCreateReq) (web.ProductCreateRes, error)
	Update(ctx context.Context, request web.ProductCreateReq) error
	Delete(ctx context.Context, CatId int) error
	// FindAll(ctx context.Context, cat *web.CatGetParam) []web.CatGetResponse
}
