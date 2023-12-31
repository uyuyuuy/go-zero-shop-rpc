package logic

import (
	"context"

	"github.com/uyuyuuy/go-zero-shop-rpc/app/cart/internal/svc"
	"github.com/uyuyuuy/go-zero-shop-rpc/app/cart/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCartLogic {
	return &DelCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelCartLogic) DelCart(in *pb.DelCartReq) (*pb.DelCartResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelCartResp{}, nil
}
