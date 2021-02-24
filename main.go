package main

import (
	"log"
	"os"

	"github.com/dinofei/gomicro-server/handlers"
	pb "github.com/dinofei/gomicro-server/proto/user"
	"github.com/dinofei/gomicro-server/trace"
	"github.com/micro/go-micro/v2"
	traceplugin "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
)

func main() {

	// 初始化全局服务追踪
	t, io, err := trace.NewTracer("go.server.user", os.Getenv("MICRO_TRACE_SERVER"))
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	service := micro.NewService(
		micro.Name("go.server.user"),
		micro.WrapHandler(traceplugin.NewHandlerWrapper(opentracing.GlobalTracer())),
	)
	service.Init()

	pb.RegisterUserServiceHandler(service.Server(), &handlers.UserServiceHandler{})

	if err := service.Run(); err != nil {
		log.Fatalln("启动失败：", err)
	}
}
