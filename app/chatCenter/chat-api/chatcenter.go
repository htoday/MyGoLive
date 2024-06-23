package main

import (
	"flag"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"
	"mygo/app/chatCenter/chat-api/dao"
	"net/http"

	"mygo/app/chatCenter/chat-api/internal/config"
	"mygo/app/chatCenter/chat-api/internal/handler"
	"mygo/app/chatCenter/chat-api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "app/chatCenter/chat-api/etc/chatcenter.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(nil, func(w http.ResponseWriter) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
	}, "*"))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		logx.Error(err, "failed to connect to RabbitMQ server")
		return
	}
	defer conn.Close()

	dao.Ch, err = conn.Channel()
	if err != nil {
		logx.Error(err, "failed to open a channel")
		return
	}
	defer dao.Ch.Close()
	err = dao.Ch.ExchangeDeclare(
		"chat",  // name
		"topic", // type
		true,    // durable
		false,   // auto-deleted
		false,   // internal
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		logx.Error(err, "failed to declare an exchange")
		return
	}
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
