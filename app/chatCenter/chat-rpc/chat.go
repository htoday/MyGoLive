package main

import (
	"flag"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"
	"mygo/app/chatCenter/chat-api/dao"

	"mygo/app/chatCenter/chat-rpc/internal/config"
	"mygo/app/chatCenter/chat-rpc/internal/server"
	"mygo/app/chatCenter/chat-rpc/internal/svc"
	"mygo/app/chatCenter/chat-rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "app/chatCenter/chat-rpc/etc/chat.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterChatRoomServiceServer(grpcServer, server.NewChatRoomServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
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
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
