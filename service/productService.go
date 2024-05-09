package service

import (
	"context"
	"eniqilo_store/model/web"
)

type ProductService interface {
	Create(ctx context.Context, request web.ProductCreateReq) web.ProductCreateRes
	Update(ctx context.Context, request web.ProductCreateReq)
	Delete(ctx context.Context, CatId int)
	// FindAll(ctx context.Context, cat *web.CatGetParam) []web.CatGetResponse
}
