package main

import (
	"austin-go/common/dbx"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"

	"austin-go/app/austin-admin/api/internal/config"
	"austin-go/app/austin-admin/api/internal/handler"
	"austin-go/app/austin-admin/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/api-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	logx.DisableStat()
	handler.RegisterHandlers(server, ctx)
	dbx.InitDb(c.Mysql)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
