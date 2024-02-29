// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/pprof"
	"github.com/yanguiyuan/cloudspace/internal/api/mw"
	"github.com/yanguiyuan/cloudspace/internal/api/router"
	"github.com/yanguiyuan/yuan/pkg/config"
	"time"
)

const identityKey = "key"

func main() {
	c := config.NewConfig()

	mw.InitJWT()
	h := server.New(server.WithHostPorts(":"+c.GetString("api.port")), server.WithMaxRequestBodySize(1024*1024*1024))
	h.Use(cors.New(cors.Config{
		AllowOrigins:     c.GetStringSlice("api.cors.allowOrigins"),
		AllowMethods:     c.GetStringSlice("api.cors.allowMethods"),
		AllowHeaders:     c.GetStringSlice("api.cors.allowHeaders"),
		ExposeHeaders:    c.GetStringSlice("api.cors.exposeHeaders"),
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	pprof.Register(h)
	hlog.Debug()

	router.Register(h)
	h.Spin()
}
