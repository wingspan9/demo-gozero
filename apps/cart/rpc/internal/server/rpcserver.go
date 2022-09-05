// Code generated by goctl. DO NOT EDIT!
// Source: rpc.proto

package server

import (
	"context"

	"demo-gozero/apps/cart/rpc/internal/logic"
	"demo-gozero/apps/cart/rpc/internal/svc"
	"demo-gozero/apps/cart/rpc/rpc"
)

type RpcServer struct {
	svcCtx *svc.ServiceContext
	rpc.UnimplementedRpcServer
}

func NewRpcServer(svcCtx *svc.ServiceContext) *RpcServer {
	return &RpcServer{
		svcCtx: svcCtx,
	}
}

func (s *RpcServer) Ping(ctx context.Context, in *rpc.Request) (*rpc.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
