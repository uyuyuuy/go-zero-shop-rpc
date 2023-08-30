package logic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/uyuyuuy/go-zero-shop-rpc/app/order/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

	var order model.Order
	err := l.svcCtx.Db.QueryRow(&order, "select * from `order` where order_id = ? for update;", in.OrderId)
	if err != nil {
		fmt.Println("orderRollback 服务异常1", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		fmt.Println("orderRollback 服务异常2")
		return nil, status.Error(codes.Internal, err.Error())
	}

	fmt.Println("orderRollback 开始执行")
	db, err := l.svcCtx.Db.RawDB()
	if dtmErr := barrier.CallWithDB(db, func(tx *sql.Tx) error {
		res, insertErr := tx.Exec("update `order` set status = 9 where order_id = ?", in.OrderId)
		if insertErr != nil {
			return fmt.Errorf("更新订单失败 err :%v", insertErr)
		}
		rows, _ := res.RowsAffected()
		if rows != 1 {
			return fmt.Errorf("更新记录数量错误，%d", rows)
		}
		return nil
	}); dtmErr != nil {
		fmt.Println("orderRollback 服务异常3", dtmErr)
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}
	fmt.Println("orderRollback 结束执行")

	return &pb.CreateOrderResp{OrderId: in.OrderId}, nil
}
