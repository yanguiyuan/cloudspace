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
<<<<<<< HEAD
		user.PUT("/:id/password", userserver.ResetPassword)
	}
	admin := root.Group("/admin", mw.JwtMiddleware.MiddlewareFunc())
	{
		admin.PUT("/user/:id/password", mw.CheckAdminPermission, userserver.AdminResetPassword)
		///list/0/3
		user.GET("/list/:offset/:limit", mw.CheckAdminPermission, userserver.GetUsers)
=======
		///list/0/3
		user.GET("/list/:offset/:limit", userserver.GetUsers)
>>>>>>> 641946dc (test)
	}
}
