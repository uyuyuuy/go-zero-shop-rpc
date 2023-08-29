package logic

import (
	"context"

	"github.com/uyuyuuy/go-zero-shop-rpc/app/order/internal/svc"
	"github.com/uyuyuuy/go-zero-shop-rpc/app/order/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderRollbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderRollbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderRollbackLogic {
	return &CreateOrderRollbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderRollbackLogic) CreateOrderRollback(in *pb.CreateOrderReq) (*pb.CreateOrderResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CreateOrderResp{}, nil
}
