package grpc

import (
	pb "my-gin-project/api"

	"go-common/library/conf/paladin"
	"go-common/library/net/rpc/warden"
)

// New new a grpc server.
func New(svc pb.DemoServer) (ws *warden.Server, err error) {
	var (
		cfg warden.ServerConfig
		ct  paladin.TOML
	)
	if err = paladin.Get("grpc.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Server").UnmarshalTOML(&cfg); err != nil {
		return
	}
	ws = warden.NewServer(&cfg)
	pb.RegisterDemoServer(ws.Server(), svc)
	ws, err = ws.Start()
	return
}
