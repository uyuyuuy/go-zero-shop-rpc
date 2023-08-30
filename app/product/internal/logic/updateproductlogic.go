package logic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dtm-labs/client/dtmcli"
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
		fmt.Println("product 服务异常1", err)
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}
	number := product.Number + in.Number

	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		fmt.Println("product 服务异常2")
		return nil, status.Error(codes.Internal, err.Error())
	}

	fmt.Println("product 开始执行")
	db, err := l.svcCtx.Db.RawDB()
	if dtmErr := barrier.CallWithDB(db, func(tx *sql.Tx) error {
		res, updateErr := tx.Exec("update product set number = ? where id = ?", number, in.ProductId)
		if updateErr != nil {
			return fmt.Errorf("更新库存失败 err :%v", updateErr)
		}
		rows, _ := res.RowsAffected()
		if rows != 1 {
			return fmt.Errorf("更新记录数量错误，%d", rows)
		}
		return nil
	}); dtmErr != nil {
		fmt.Println("product 服务异常3")
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}
	fmt.Println("product 结束执行")
	return &pb.UpdateResp{Number: number}, nil
}
