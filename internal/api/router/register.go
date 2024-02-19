package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/yanguiyuan/cloudspace/internal/api/handler"
<<<<<<< HEAD
	fileserver "github.com/yanguiyuan/cloudspace/internal/api/handler/file"
=======
>>>>>>> 641946dc (test)
	"github.com/yanguiyuan/cloudspace/internal/api/mw"
	"github.com/yanguiyuan/cloudspace/internal/api/router/file"
	"github.com/yanguiyuan/cloudspace/internal/api/router/user"
)

func Register(r *server.Hertz) {
	r.GET("/", handler.Ping)
	r.POST("/login", mw.JwtMiddleware.LoginHandler)
<<<<<<< HEAD
	r.GET("/check", fileserver.CheckLock)
=======

>>>>>>> 641946dc (test)
	user.Register(r)
	file.Register(r)
}
