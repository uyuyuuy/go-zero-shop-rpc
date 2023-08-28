package logic

import (
	"context"

	"github.com/uyuyuuy/go-zero-shop-rpc/app/cart/internal/svc"
	"github.com/uyuyuuy/go-zero-shop-rpc/app/cart/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCartLogic {
	return &AddCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCartLogic) AddCart(in *pb.AddCartReq) (*pb.AddCartResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddCartResp{}, nil
}
