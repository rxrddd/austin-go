package main

import (
	"austin-go/app/austin-job/internal/config"
	"austin-go/app/austin-job/internal/handler/handlers"
	"austin-go/app/austin-job/internal/listen"
	"austin-go/app/austin-job/internal/svc"
	"austin-go/common/dbx"
	"flag"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/austin-job.yaml", "the config file")

func main() {
	flag.Parse()
	logx.DisableStat()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	group := service.NewServiceGroup()
	for _, mq := range listen.Mqs(ctx) {
		group.Add(mq)
	}

	handlers.SetUp(ctx)
	dbx.InitDb(c.Mysql)

	defer group.Stop()
	group.Start()

}
