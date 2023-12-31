// Code generated by goctl. DO NOT EDIT.
// Source: cart.proto

package cart

import (
	"context"

	"github.com/uyuyuuy/go-zero-shop-rpc/app/cart/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddCartReq     = pb.AddCartReq
	AddCartResp    = pb.AddCartResp
	Cart           = pb.Cart
	CartListReq    = pb.CartListReq
	CartListResp   = pb.CartListResp
	DelCartReq     = pb.DelCartReq
	DelCartResp    = pb.DelCartResp
	UpdateCartReq  = pb.UpdateCartReq
	UpdateCartResp = pb.UpdateCartResp

	CartZrpcClient interface {
		CartList(ctx context.Context, in *CartListReq, opts ...grpc.CallOption) (*CartListResp, error)
		AddCart(ctx context.Context, in *AddCartReq, opts ...grpc.CallOption) (*AddCartResp, error)
		UpdateCart(ctx context.Context, in *UpdateCartReq, opts ...grpc.CallOption) (*UpdateCartResp, error)
		DelCart(ctx context.Context, in *DelCartReq, opts ...grpc.CallOption) (*DelCartResp, error)
	}

	defaultCartZrpcClient struct {
		cli zrpc.Client
	}
)

func NewCartZrpcClient(cli zrpc.Client) CartZrpcClient {
	return &defaultCartZrpcClient{
		cli: cli,
	}
}

func (m *defaultCartZrpcClient) CartList(ctx context.Context, in *CartListReq, opts ...grpc.CallOption) (*CartListResp, error) {
	client := pb.NewCartClient(m.cli.Conn())
	return client.CartList(ctx, in, opts...)
}

func (m *defaultCartZrpcClient) AddCart(ctx context.Context, in *AddCartReq, opts ...grpc.CallOption) (*AddCartResp, error) {
	client := pb.NewCartClient(m.cli.Conn())
	return client.AddCart(ctx, in, opts...)
}

func (m *defaultCartZrpcClient) UpdateCart(ctx context.Context, in *UpdateCartReq, opts ...grpc.CallOption) (*UpdateCartResp, error) {
	client := pb.NewCartClient(m.cli.Conn())
	return client.UpdateCart(ctx, in, opts...)
}

func (m *defaultCartZrpcClient) DelCart(ctx context.Context, in *DelCartReq, opts ...grpc.CallOption) (*DelCartResp, error) {
	client := pb.NewCartClient(m.cli.Conn())
	return client.DelCart(ctx, in, opts...)
}
