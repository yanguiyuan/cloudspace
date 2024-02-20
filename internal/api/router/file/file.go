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
			file.GET("/:id/download", mw.CheckReadPermission, fileserver.Download)
			file.GET("/:id/check", fileserver.CheckFileLock)
			//查询文件项根据id
			file.GET("/:id", mw.CheckReadPermission, fileserver.QueryFileItemByID)
			file.GET("/:id/content", mw.CheckReadPermission, fileserver.FetchFileContent)
			//查询文件项列表，每个文件项的父文件项ID为id
			file.GET("/:id/list", mw.CheckReadPermission, fileserver.QueryFileItemList)
			file.POST("/:id/upload", fileserver.Upload)
			file.DELETE("/:id", mw.CheckWritePermission, fileserver.DeleteFileItem)
			file.GET("/:id/url", mw.CheckReadPermission, fileserver.GetFileURL)
			file.PUT("/:id/rename", mw.CheckWritePermission, fileserver.Rename)
			file.PUT("/:id/content", mw.CheckWritePermission, fileserver.ModifyFileContent)
			file.POST("/:id/content", mw.CheckWritePermission, fileserver.CreateTextFile)
		}
		directory := user.Group("/directory")
		{
			directory.POST("", mw.CheckWritePermission, fileserver.CreateDirectory)
			directory.DELETE("/:id", mw.CheckWritePermission, fileserver.RemoveDirectory)
		}
		namespace := user.Group("/namespace")
		{
			namespace.GET("/:id", fileserver.QueryNamespace)
			namespace.GET("/list", fileserver.QueryUserNamespaces)
			namespace.POST("/:name", fileserver.CreateNamespace)
			namespace.GET("/:id/link", mw.CheckAllPermission, fileserver.NamespaceLink)
			namespace.POST("/:id/link", fileserver.LinkNamespace)
		}
	}

}
