package logic

import (
	"context"

	"github.com/uyuyuuy/go-zero-shop-rpc/app/product/internal/svc"
	"github.com/uyuyuuy/go-zero-shop-rpc/app/product/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductRollbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductRollbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductRollbackLogic {
	return &UpdateProductRollbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProductRollbackLogic) UpdateProductRollback(in *pb.UpdateProductReq) (*pb.UpdateResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateResp{}, nil
}
