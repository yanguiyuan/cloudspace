package user

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/yanguiyuan/cloudspace/internal/api/handler/user"
)

func Register(r *server.Hertz) {
	root := r.Group("/")
	root.POST("/register", user.Register)
}
