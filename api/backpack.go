package main

import (
	"UBC/api/library/xerr"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"

	"UBC/api/internal/config"
	"UBC/api/internal/handler"
	"UBC/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "D:\\goproject\\UBC\\api\\etc\\backpack.yaml", "the config file")

//var configFile = flag.String("f", "D:\\goproject\\UBC\\api\\etc\\backpack.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	//server := rest.MustNewServer(c.RestConf, rest.WithCors())
	server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(cors, nil, "*"))

	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	xerr.Init("zh")
	logx.DisableStat()
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func cors(header http.Header) {
	header.Set("Access-Control-Allow-Headers", "*")
	header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
	header.Set("Access-Control-Expose-Headers", "Content-Length, Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
	header.Set("Access-Control-Allow-Credentials", "true")
}
