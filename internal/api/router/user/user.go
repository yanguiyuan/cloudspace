package user

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	userserver "github.com/yanguiyuan/cloudspace/internal/api/handler/user"
	"github.com/yanguiyuan/cloudspace/internal/api/mw"
)

func Register(r *server.Hertz) {
	root := r.Group("/")
	root.POST("/register", userserver.Register)
	user := root.Group("/user", mw.JwtMiddleware.MiddlewareFunc())
	{
		user.POST("/logout", mw.JwtMiddleware.LogoutHandler)
		user.GET("/info", userserver.UserInfo)
		user.PUT("/info", userserver.ModifyUserInfo)
		user.PUT("/:id/password", userserver.ResetPassword)
	}
	admin := root.Group("/admin", mw.JwtMiddleware.MiddlewareFunc())
	{
		admin.GET("/list/:offset/:limit", userserver.GetUsers)
		admin.PUT("/:id/password", userserver.AdminResetPassword)
	}
}
