package logic

import (
	"context"

	"github.com/uyuyuuy/go-zero-shop-rpc/app/cart/internal/svc"
	"github.com/uyuyuuy/go-zero-shop-rpc/app/cart/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCartListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartListLogic {
	return &CartListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CartListLogic) CartList(in *pb.CartListReq) (*pb.CartListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CartListResp{}, nil
}
