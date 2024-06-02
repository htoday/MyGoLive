package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"log"
	"mygo/app/userCenter/user-api/internal/config"
	"mygo/app/userCenter/user-api/internal/handler"
	"mygo/app/userCenter/user-api/internal/svc"
)

var configFile = flag.String("f", "app/userCenter/user-api/etc/usercenter.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	err := c.SetUp()
	if err != nil {
		log.Fatal(err)
	}
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	//fs := http.FileServer(http.Dir("./path/to/your/static/files"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
