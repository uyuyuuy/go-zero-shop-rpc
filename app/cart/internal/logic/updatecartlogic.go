package logic

import (
	"context"

	"github.com/uyuyuuy/go-zero-shop-rpc/app/cart/internal/svc"
	"github.com/uyuyuuy/go-zero-shop-rpc/app/cart/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCartLogic {
	return &UpdateCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCartLogic) UpdateCart(in *pb.UpdateCartReq) (*pb.UpdateCartResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateCartResp{}, nil
}
