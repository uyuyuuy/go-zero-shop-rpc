// Code generated by goctl. DO NOT EDIT.
// Source: product.proto

package server

import (
	"context"

	"github.com/uyuyuuy/go-zero-shop-rpc/app/product/internal/logic"
	"github.com/uyuyuuy/go-zero-shop-rpc/app/product/internal/svc"
	"github.com/uyuyuuy/go-zero-shop-rpc/app/product/pb"
)

type ProductServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedProductServer
}

func NewProductServer(svcCtx *svc.ServiceContext) *ProductServer {
	return &ProductServer{
		svcCtx: svcCtx,
	}
}

func (s *ProductServer) UpdateProduct(ctx context.Context, in *pb.UpdateProductReq) (*pb.UpdateResp, error) {
	l := logic.NewUpdateProductLogic(ctx, s.svcCtx)
	return l.UpdateProduct(in)
}

func (s *ProductServer) UpdateProductRollback(ctx context.Context, in *pb.UpdateProductReq) (*pb.UpdateResp, error) {
	l := logic.NewUpdateProductRollbackLogic(ctx, s.svcCtx)
	return l.UpdateProductRollback(in)
}
