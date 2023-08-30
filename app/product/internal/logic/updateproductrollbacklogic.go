package logic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/uyuyuuy/go-zero-shop-rpc/app/product/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
	var product model.Product
	err := l.svcCtx.Db.QueryRow(&product, "select * from product where id = ? for update;", in.ProductId)
	if err != nil {
		fmt.Println("productRollback 服务异常1", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	number := product.Number - in.Number

	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		fmt.Println("productRollback 服务异常2", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	fmt.Println("productRollback 开始执行", err)

	db, err := l.svcCtx.Db.RawDB()
	if dtmErr := barrier.CallWithDB(db, func(tx *sql.Tx) error {
		fmt.Println("productRollback 执行回滚")
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
		fmt.Println("productRollback 服务异常3", dtmErr)
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}

	fmt.Println("productRollback 结束执行")

	return &pb.UpdateResp{Number: number}, nil
}
