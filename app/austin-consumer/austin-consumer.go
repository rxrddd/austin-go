package main

import (
	"austin-go/app/austin-consumer/internal/listen"
	"flag"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"

	"austin-go/app/austin-consumer/internal/config"
	"austin-go/app/austin-consumer/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/austin-consumer.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	group := service.NewServiceGroup()
	for _, mq := range listen.Mqs(ctx) {
		group.Add(mq)
	}

	logx.DisableStat()
	defer group.Stop()
	group.Start()
}
