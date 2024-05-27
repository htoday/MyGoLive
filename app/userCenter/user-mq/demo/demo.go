package main

import (
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/service"
	"mygo/app/userCenter/user-mq/demo/internal/mqs"

	"mygo/app/userCenter/user-mq/demo/internal/config"
	"mygo/app/userCenter/user-mq/demo/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "app/userCenter/user-mq/demo/etc/demo-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	svcCtx := svc.NewServiceContext(c)
	ctx := context.Background()
	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	for _, mq := range mqs.Consumers(c, ctx, svcCtx) {
		serviceGroup.Add(mq)
	}

	serviceGroup.Start()
}
