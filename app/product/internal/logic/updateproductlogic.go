package logic

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/uyuyuuy/go-zero-shop-rpc/app/product/internal/model"
	"github.com/uyuyuuy/go-zero-shop-rpc/app/product/internal/svc"
	"github.com/uyuyuuy/go-zero-shop-rpc/app/product/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	barrier.CallWithDB(l.svcCtx.Db.RawDB(), func(tx *sql.Tx) error {

	})
	_, updateErr := l.svcCtx.Db.Exec("update product set number = ?", number)
	if updateErr != nil {
		return &pb.UpdateResp{}, updateErr
	}
	return &pb.UpdateResp{Number: number}, nil
}
