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
	"net/http"
)

var configFile = flag.String("f", "app/userCenter/user-api/etc/usercenter.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	//server := rest.MustNewServer(c.RestConf)
	server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(nil, func(w http.ResponseWriter) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
	}, "*"))

	defer server.Stop()
	err := c.SetUp()
	if err != nil {
		log.Fatal(err)
	}
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	go func() {
		fs := http.FileServer(http.Dir("frontend/dist"))
		http.Handle("/", http.StripPrefix("/", fs))
		http.ListenAndServe(":8083", nil)
	}()
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
	//handler.StaticFileHandler(server, ctx)

}
