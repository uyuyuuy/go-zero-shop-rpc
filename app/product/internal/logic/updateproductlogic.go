package logic

import (
	"context"
	"github.com/uyuyuuy/go-zero-shop-rpc/app/product/internal/model"
	"github.com/uyuyuuy/go-zero-shop-rpc/app/product/internal/svc"
	"github.com/uyuyuuy/go-zero-shop-rpc/app/product/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductLogic {
	return &UpdateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProductLogic) UpdateProduct(in *pb.UpdateProductReq) (*pb.UpdateResp, error) {
	// todo: add your logic here and delete this line
	var product model.Product
	err := l.svcCtx.Db.QueryRow(&product, "select * from product where id = ? for update;", in.ProductId)
	if err != nil {
		return &pb.UpdateResp{}, err
	}
	number := product.Number + in.Number
	_, updateErr := l.svcCtx.Db.Exec("update product set number = ?", number)
	if updateErr != nil {
		return &pb.UpdateResp{}, updateErr
	}
	return &pb.UpdateResp{Number: number}, nil
}
