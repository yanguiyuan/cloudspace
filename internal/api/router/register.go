package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/yanguiyuan/cloudspace/internal/api/handler"
	"github.com/yanguiyuan/cloudspace/internal/api/mw"
	"github.com/yanguiyuan/cloudspace/internal/api/router/file"
	"github.com/yanguiyuan/cloudspace/internal/api/router/user"
)

func Register(r *server.Hertz) {
	r.GET("/", handler.Ping)
	r.POST("/login", mw.JwtMiddleware.LoginHandler)

	user.Register(r)
	file.Register(r)
}
