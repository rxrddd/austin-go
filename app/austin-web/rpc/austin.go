package main

import (
	"austin-go/common/dbx"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"

	"austin-go/app/austin-web/rpc/austin"
	"austin-go/app/austin-web/rpc/internal/config"
	"austin-go/app/austin-web/rpc/internal/server"
	"austin-go/app/austin-web/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/austin.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewAustinServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		austin.RegisterAustinServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	dbx.InitDb(c.Mysql)

	logx.DisableStat()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
