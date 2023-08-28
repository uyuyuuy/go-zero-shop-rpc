package main

import (
	"flag"
	"fmt"

	"github.com/uyuyuuy/go-zero-shop-rpc/app/cart/internal/config"
	"github.com/uyuyuuy/go-zero-shop-rpc/app/cart/internal/server"
	"github.com/uyuyuuy/go-zero-shop-rpc/app/cart/internal/svc"
	"github.com/uyuyuuy/go-zero-shop-rpc/app/cart/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/cart.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterCartServer(grpcServer, server.NewCartServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
