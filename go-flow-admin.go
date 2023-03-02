package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/sirupsen/logrus"
	"go-flow-admin/config"
	"go-flow-admin/middleware"
	"go-flow-admin/model"
	"go-flow-admin/pkg/logger"
	"go-flow-admin/pkg/swagger"
	"go-flow-admin/query"
	"go-flow-admin/routes"

	"github.com/gin-gonic/gin"
)

//go:generate goctl api plugin -p gengin -api go-flow-admin.api -dir .
//go:generate goctl api plugin -p "goctl-swagger swagger -filename asset/swagger/swagger.json" -api go-flow-admin.api -dir .
func main() {
	flag.Parse()

	configFile := fmt.Sprintf("etc/go-flow-admin-%s.yaml", config.Env)

	e := gin.New()
	e.Use(
		middleware.RequestId,
		middleware.RequestLog,
		gin.Recovery(),
		middleware.CorsMiddleware,
	)

	config.Setup(configFile)
	logger.Setup()
	flowablesdk.Setup(flowablesdk.Config{Url: "http://127.0.0.1:8080/", RequestDebug: true, ResponseDebug: true})
	model.Setup()
	query.SetDefault(model.DB())
	routes.Setup(e)
	swagger.Setup(e)

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.ServerConf.Host, config.ServerConf.Port),
		Handler: e,
	}
	go func() {
		logrus.Info("listen to ", server.Addr)
		server.ListenAndServe()
	}()
	wait := config.RegisterCloseFn(func() {
		defer logrus.Warning("closed api server")

		ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancelFunc()
		server.Shutdown(ctx)
	})
	wait()
}
