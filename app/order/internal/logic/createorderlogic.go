package logic

import (
	"context"
	"fmt"

	"github.com/uyuyuuy/go-zero-shop-rpc/app/order/internal/svc"
	"github.com/uyuyuuy/go-zero-shop-rpc/app/order/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderLogic) CreateOrder(in *pb.CreateOrderReq) (*pb.CreateOrderResp, error) {
	// todo: add your logic here and delete this line

	r, err := l.svcCtx.Db.Exec("insert into order(order_id,user_id,product_id,number) values(?,?,?,?)", in.OrderId, in.UserId, in.ProductId, in.Number)
	if err != nil {
		return &pb.CreateOrderResp{}, err
	}
	fmt.Println(r.LastInsertId())

	return &pb.CreateOrderResp{OrderId: in.OrderId}, nil
}
