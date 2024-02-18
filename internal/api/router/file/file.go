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
			//查询文件项根据id
			file.GET("/:id", fileserver.QueryFileItemByID)
			file.GET("/:id/content", fileserver.FetchFileContent)
			//查询文件项列表，每个文件项的父文件项ID为id
			file.GET("/:id/list", mw.CheckReadPermission, fileserver.QueryFileItemList)
			file.POST("/:id/upload", fileserver.Upload)
			file.DELETE("/:id", fileserver.DeleteFileItem)
			file.GET("/:id/url", fileserver.GetFileURL)
			file.PUT("/rename", fileserver.Rename)
			file.PUT("/:id/content", fileserver.ModifyFileContent)
			file.POST("/:parent_id/content", fileserver.CreateTextFile)
		}
		directory := user.Group("/directory")
		{
			directory.POST("", fileserver.CreateDirectory)
			directory.DELETE("/:id", fileserver.RemoveDirectory)
		}
		namespace := user.Group("/namespace")
		{
			namespace.GET("/:id", fileserver.QueryNamespace)
			namespace.GET("/list", fileserver.QueryUserNamespaces)
			namespace.POST("/:name", fileserver.CreateNamespace)
			namespace.GET("/:id/link", fileserver.NamespaceLink)
			namespace.POST("/:id/link", fileserver.LinkNamespace)
		}
	}

}
