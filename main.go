package main

import (
	"flag"
	"fmt"
	"github.com/ZUCCzwp/kitex_ddd_example/internal/config"
	"github.com/ZUCCzwp/kitex_ddd_example/internal/interfaces"
	"github.com/ZUCCzwp/kitex_ddd_example/internal/svc"
	"github.com/ZUCCzwp/kitex_ddd_example/kitex_gen/hello/helloservice"
	"github.com/ZUCCzwp/kitex_ddd_example/pkg/conf"
	"github.com/ZUCCzwp/kitex_ddd_example/pkg/onelog"
	kitexlogrus "github.com/ZUCCzwp/kitex_ddd_example/pkg/onelogrus"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"net"
	"os"
)

var configFile = flag.String("f", "etc/config.yaml", "the config file")

func main() {
	var c config.Config
	conf.MustLoad(*configFile, &c)

	onelog.SetLogger(kitexlogrus.NewLogger())
	onelog.SetLevel(onelog.LevelDebug)

	svr := server.NewServer(server.WithServiceAddr(&net.TCPAddr{Port: c.Port}),
		server.WithSuite(tracing.NewServerSuite(tracing.WithEnableMetadata())),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: c.Name}))

	//p := provider.NewOpenTelemetryProvider(
	//	provider.WithServiceName(os.Getenv("OTEL_SERVICE_NAME")),
	//	//provider.WithExportEndpoint("192.168.105.196:32167"),
	//	provider.WithInsecure(),
	//)
	//defer p.Shutdown(context.Background())
	ctx := svc.NewServiceContext(c)
	// TODO: register your services here
	err := helloservice.RegisterService(svr, &interfaces.HelloServiceImpl{
		SvcCtx: ctx,
	})
	if err != nil {
		return
	}

	err = svr.Run()
	if err != nil {
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			os.Exit(1)
		}
	}()
}
