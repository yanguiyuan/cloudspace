package file

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/yanguiyuan/cloudspace/internal/api/mw"
)
import fileserver "github.com/yanguiyuan/cloudspace/internal/api/handler/file"

func Register(r *server.Hertz) {
	root := r.Group("/")
	user := root.Group("/user", mw.JwtMiddleware.MiddlewareFunc())
	{
		file := user.Group("/file")
		{
			//查询用户的默认命名空间的根文件项（namespace文件项）
			file.GET("/root", fileserver.UserRootFile)
			//查询文件项根据id
			file.GET("/:id", fileserver.QueryFileItemByID)
			//查询文件项列表，每个文件项的父文件项ID为id
			file.GET("/:id/list", mw.CheckReadPermission, fileserver.QueryFileItemList)
			file.POST("/:id/upload", fileserver.Upload)
			file.DELETE("/:id/:name", fileserver.DeleteFileItem)
			file.GET("/:id/url", fileserver.GetFileURL)
			file.PUT("/rename", fileserver.Rename)
		}
		directory := user.Group("/directory")
		{
			directory.POST("/", fileserver.CreateDirectory)
		}
		namespace := user.Group("/namespace")
		{
			namespace.GET("/list", fileserver.QueryUserNamespaces)
			namespace.POST("/:name", fileserver.CreateNamespace)
			namespace.GET("/:id/link", fileserver.LinkNamespace)
		}
	}

}
