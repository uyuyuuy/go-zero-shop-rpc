package logic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

	//var o model.Order
	//err := l.svcCtx.Db.QueryRow(&o, "select * from `order` where order_id = ?", in.OrderId)
	//if err != nil || o.OrderId != "" {
	//	fmt.Println("order 服务异常1", err)
	//	return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	//}

	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		fmt.Println("order 服务异常2")
		return nil, status.Error(codes.Internal, err.Error())
	}

	fmt.Println("order 开始执行")
	db, err := l.svcCtx.Db.RawDB()
	if dtmErr := barrier.CallWithDB(db, func(tx *sql.Tx) error {
		res, insertErr := tx.Exec("insert into `order`(order_id,user_id,product_id,number) values(?,?,?,?)", in.OrderId, in.UserId, in.ProductId, in.Number)
		if insertErr != nil {
			return fmt.Errorf("更新库存失败 err :%v", insertErr)
		}
		rows, _ := res.RowsAffected()
		if rows != 1 {
			return fmt.Errorf("更新记录数量错误，%d", rows)
		}
		return nil
	}); dtmErr != nil {
		fmt.Println("order 服务异常3", dtmErr)
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}
	fmt.Println("order 结束执行")

	return &pb.CreateOrderResp{OrderId: in.OrderId}, nil
}
