package grpc

import (
	pb "github.com/znyh/middle-end/proto/logserver"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
)

// New new a grpc server.
func New(svc pb.LogserverServer) (ws *warden.Server, err error) {
	var cfg struct {
		Server *warden.ServerConfig
	}
	if err = paladin.Get("grpc.txt").UnmarshalTOML(&cfg); err != nil {
		return
	}
	ws = warden.NewServer(cfg.Server)
	pb.RegisterLogserverServer(ws.Server(), svc)
	ws, err = ws.Start()
	return
}
