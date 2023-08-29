package svc

import (
	"github.com/uyuyuuy/go-zero-shop-rpc/app/order/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Db     sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.MysqlDsn)
	return &ServiceContext{
		Config: c,
		Db:     conn,
	}
}
