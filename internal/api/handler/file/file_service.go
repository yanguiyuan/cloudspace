package file

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/golang-jwt/jwt/v4"
	"github.com/muesli/cache2go"
	"github.com/yanguiyuan/cloudspace/internal/api/handler"
	"github.com/yanguiyuan/cloudspace/internal/api/model/file"
	"github.com/yanguiyuan/cloudspace/internal/api/mw"
	"github.com/yanguiyuan/cloudspace/internal/cloudfile"
	"github.com/yanguiyuan/cloudspace/pkg/errno"
	"github.com/yanguiyuan/cloudspace/pkg/rpc"
	"github.com/yanguiyuan/yuan/pkg/config"
	"io"
	"strconv"
	"time"
)

var cache = cache2go.Cache("fileWrite")

// Query .
// @router /files/:parent_id [GET]
func QueryFileItemList(ctx context.Context, c *app.RequestContext) {
	parentID := c.Param("id")
	client, err := cloudfile.NewFileServiceClient()
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	identity, b := c.Get(mw.IdentityKey)
	if !b {
		c.String(consts.StatusUnauthorized, "未登录")
	}
	r, err := client.Query(ctx, parentID, identity.(int64))
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.NotFoundCode,
			"message": errno.NotFoundMsg,
		})
		fmt.Println(err)
		return
	}
	c.JSON(consts.StatusOK, utils.H{
		"code":    errno.SuccessCode,
		"message": errno.SuccessMsg,
		"data":    r,
	})
}

// Upload .
// @router /file/upload [POST]
func Upload(ctx context.Context, c *app.RequestContext) {
	form, _ := c.MultipartForm()
	files := form.File["data"]
	parentID := c.Param("id")
	client, err := cloudfile.NewFileServiceClient()
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	var item *rpc.CloudFileItem
	for _, file := range files {
		fmt.Println(file.Filename)
		f, err := file.Open()
		if err != nil {
			c.String(consts.StatusInternalServerError, err.Error())
			return
		}
		d, err := io.ReadAll(f)
		if err != nil {
			c.String(consts.StatusInternalServerError, err.Error())
			return
		}
		it, err := client.QueryFileItemByID(ctx, parentID)
		if err != nil {
			c.JSON(consts.StatusOK, utils.H{
				"code":    errno.ServiceErrCode,
				"message": err.Error(),
			})
			return
		}
		item, err = client.UploadFile(ctx, &rpc.UploadFileRequest{
			FileData:    d,
			FileName:    file.Filename,
			ParentID:    parentID,
			NamespaceID: it.NamespaceID,
		})
		if err != nil {
			c.JSON(consts.StatusOK, utils.H{
				"code":    errno.ServiceErrCode,
				"message": err.Error(),
			})
			return
		}
	}
	c.JSON(consts.StatusOK, utils.H{
		"code":    errno.SuccessCode,
		"message": errno.SuccessMsg,
		"data":    item,
	})
}

func QueryFileItemByID(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")
	client, err := cloudfile.NewFileServiceClient()
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	byID, err := client.QueryFileItemByID(ctx, id)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.NotFoundCode,
			"message": errno.NotFoundMsg,
		})
		return
	}
	c.JSON(consts.StatusOK, utils.H{
		"code":    errno.SuccessCode,
		"message": errno.SuccessMsg,
		"data":    byID,
	})
}
func CreateDirectory(ctx context.Context, c *app.RequestContext) {
	var req rpc.CreateDirectoryRequest
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.InvalidParamCode,
			"message": errno.InvalidParamMsg,
		})
		return
	}
	client, err := cloudfile.NewFileServiceClient()
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	res, err := client.CreateDirectory(ctx, &req)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.ServiceErrCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(consts.StatusOK, utils.H{
		"code":    errno.SuccessCode,
		"message": errno.SuccessMsg,
		"data":    res,
	})
}
func GetFileURL(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")
	var resp string
	client, err := cloudfile.NewFileServiceClient()
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	resp, err = client.GetFileURL(ctx, id)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.NotFoundCode,
			"message": errno.NotFoundMsg,
		})
		return
	}

	c.JSON(consts.StatusOK, utils.H{
		"code":    errno.SuccessCode,
		"message": errno.SuccessMsg,
		"data":    resp,
	})
}
func DeleteFileItem(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")
	client, err := cloudfile.NewFileServiceClient()
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	err = client.Remove(ctx, id)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.ServiceErrCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(consts.StatusOK, utils.H{
		"code":    errno.SuccessCode,
		"message": errno.SuccessMsg,
	})
}
func Rename(ctx context.Context, c *app.RequestContext) {
	var req file.RenameReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.InvalidParamCode,
			"message": errno.InvalidParamMsg,
		})
		return
	}
	client, err := cloudfile.NewFileServiceClient()
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	err = client.Rename(ctx, req.ID, req.NewName)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.ServiceErrCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(consts.StatusOK, utils.H{
		"code":    errno.SuccessCode,
		"message": errno.SuccessMsg,
	})
}

func QueryUserNamespaces(ctx context.Context, c *app.RequestContext) {
	identity, b := c.Get(mw.IdentityKey) //*
	if !b {
		c.String(consts.StatusUnauthorized, "未登录")
	}
	client, err := cloudfile.NewFileServiceClient()
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	resp, err := client.QueryUserNamespaces(ctx, identity.(int64)) //*
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.ServiceErrCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(consts.StatusOK, utils.H{
		"code":    errno.SuccessCode,
		"message": errno.SuccessMsg,
		"data":    resp,
	})
}

func CreateNamespace(ctx context.Context, c *app.RequestContext) {
	name := c.Param("name")
	userid, b := handler.ExtractID(c)
	if !b {
		return
	}
	client, err := cloudfile.NewFileServiceClient()
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	namespaceID, err := client.CreateNamespace(ctx, name)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.ServiceErrCode,
			"message": err.Error(),
		})
		return
	}
	_, err = client.CreateFileItem(ctx, name, "namespace", "@root", namespaceID)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.ServiceErrCode,
			"message": err.Error(),
		})
		return
	}

	err = client.CreateUserNamespace(ctx, userid, namespaceID, 1)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.ServiceErrCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(consts.StatusOK, utils.H{
		"code":    errno.SuccessCode,
		"message": errno.SuccessMsg,
		"data":    namespaceID,
	})
}

func NamespaceLink(ctx context.Context, c *app.RequestContext) {
	//c.Redirect(consts.StatusOK, []byte("http://localhost:5173/link?user_id="))
	id, b := handler.ExtractID(c)
	namespaceId := c.Param("id")
	authority := c.Query("authority")
	name := c.Query("name")
	if !b {
		return
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["namespace_id"] = namespaceId
	claims["authority"] = authority
	claims["uid"] = id
	signedString, err := token.SignedString([]byte(config.GetString("api.jwt.secret")))
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.InvalidParamCode,
			"message": "Token签名失败:" + err.Error(),
		})
		return
	}
	c.JSON(consts.StatusOK, utils.H{
		"code":    errno.SuccessCode,
		"message": errno.SuccessMsg,
		"data": "http://localhost:5173/link?user_id=" +
			strconv.FormatInt(id, 10) +
			"&namespace_id=" + namespaceId +
			"&name=" + name +
			"&authority=" + authority +
			"&token=" + signedString +
			"&time=" + strconv.FormatInt(time.Now().Unix(), 10),
	})
}

func QueryNamespace(ctx context.Context, c *app.RequestContext) {

}

func LinkNamespace(ctx context.Context, c *app.RequestContext) {
	userId, b := handler.ExtractID(c)
	if !b {
		return
	}
	namespaceId := c.Param("id")
	authority := c.Query("authority")
	tokenString := c.Query("token")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetString("api.jwt.secret")), nil
	})
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.InvalidParamCode,
			"message": "Token解析失败",
		})
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	if claims["namespace_id"] != namespaceId {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.InvalidParamCode,
			"message": "Namespace_id与对应Token不一致，前者是" + namespaceId + ",后者是",
		})
		return
	}

	if claims["authority"] != authority {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.InvalidParamCode,
			"message": "authority与对应Token不一致",
		})
		return
	}
	client, err := cloudfile.NewFileServiceClient()
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	parseInt, err := strconv.ParseInt(namespaceId, 10, 64)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.InvalidParamCode,
			"message": "解析namespaceID失败",
		})
		return
	}
	i, err := strconv.ParseInt(authority, 10, 32)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.InvalidParamCode,
			"message": "解析authority失败",
		})
		return
	}
	err = client.LinkNamespace(ctx, parseInt, userId, int32(i))
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.ServiceErrCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(consts.StatusOK, utils.H{
		"code":    errno.SuccessCode,
		"message": errno.SuccessMsg,
	})

}

func RemoveDirectory(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")

	client, err := cloudfile.NewFileServiceClient()
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	err = client.RemoveDirectory(ctx, id)
	if err != nil {
		if err.Error() == "remote or network error[remote]: biz error: 目录不为空，无法进行删除" {
			c.JSON(consts.StatusOK, utils.H{
				"code":    errno.FolderNotEmptyCode,
				"message": "目录不为空，无法进行删除",
			})
			return
		}
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.ServiceErrCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(consts.StatusOK, utils.H{
		"code":    errno.SuccessCode,
		"message": errno.SuccessMsg,
	})
}

func FetchFileContent(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")
	client, err := cloudfile.NewFileServiceClient()
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	data, err := client.FetchFileData(ctx, id)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.ServiceErrCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(consts.StatusOK, utils.H{
		"code":    errno.SuccessCode,
		"message": errno.SuccessMsg,
		"data":    string(data),
	})
}

func ModifyFileContent(ctx context.Context, c *app.RequestContext) {
	var req file.ModifyFileContentReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.InvalidParamCode,
			"message": err.Error(),
		})
		return
	}
	client, err := cloudfile.NewFileServiceClient()
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	err = client.ModifyFileContent(ctx, req.ID, req.Content)
	if err != nil {
		fmt.Println(err)
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.ServiceErrCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(consts.StatusOK, utils.H{
		"code":    errno.SuccessCode,
		"message": errno.SuccessMsg,
	})

}

func CreateTextFile(ctx context.Context, c *app.RequestContext) {
	var req file.CreateTextFileReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.InvalidParamCode,
			"message": err.Error(),
		})
	}
	client, err := cloudfile.NewFileServiceClient()
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	r, err := client.QueryFileItemByID(ctx, req.ParentID)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.ServiceErrCode,
			"message": err.Error(),
		})
		return
	}
	resp, err := client.CreateTextFile(ctx, req.FileName, req.ParentID, req.Content, r.NamespaceID)
	if err != nil {
		fmt.Println(err)
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.ServiceErrCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(consts.StatusOK, utils.H{
		"code":    errno.SuccessCode,
		"message": errno.SuccessMsg,
		"data":    resp,
	})
}

func CheckFileLock(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")
	username := c.Query("username")
	if cache.Exists(id) {
		value, err := cache.Value(id)
		if err != nil {
			return
		}
		if username == value.Data().(string) {
			_, err := cache.Delete(id)
			if err != nil {
				fmt.Println(err)
				return
			}
			c.JSON(consts.StatusOK, utils.H{
				"code":    errno.SuccessCode,
				"message": id + "已解锁",
				"data":    "unlock",
			})
			return
		}
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.SuccessCode,
			"message": value.Data().(string) + "正在修改中，请稍后重试",
			"data":    "locked",
		})
		return
	}
	cache.Add(id, 0, username)
	c.JSON(consts.StatusOK, utils.H{
		"code":    errno.SuccessCode,
		"message": "文件成功被锁定",
		"data":    "lock",
	})
}
func CheckLock(ctx context.Context, c *app.RequestContext) {
	cache.Foreach(func(key interface{}, item *cache2go.CacheItem) {
		fmt.Println(key, item.Data())
	})
	c.JSON(consts.StatusOK, utils.H{
		"code":    errno.SuccessCode,
		"message": errno.SuccessMsg,
		"data":    cache,
	})
}

func Download(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")
	client, err := cloudfile.NewFileServiceClient()
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	data, err := client.FetchFileData(ctx, id)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.ServiceErrCode,
			"message": err.Error(),
		})
		return
	}

	it, err := client.QueryFileItemByID(ctx, id)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.ServiceErrCode,
			"message": err.Error(),
		})
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+it.FileName)
	c.Data(consts.StatusOK, "application/octet-stream", data)
}

func QueryNamespaceUsers(ctx context.Context, c *app.RequestContext) {
	namespaceIDStr := c.Param("id")
	client, err := cloudfile.NewFileServiceClient()
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	namespaceID, err := strconv.ParseInt(namespaceIDStr, 10, 64)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.ServiceErrCode,
			"message": err.Error(),
		})
		return
	}
	resp, err := client.QueryNamespaceUsers(ctx, namespaceID)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.ServiceErrCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(consts.StatusOK, utils.H{
		"code":    errno.SuccessCode,
		"message": errno.SuccessMsg,
		"data":    resp,
	})
}

func RemoveNamespaceAuthority(ctx context.Context, c *app.RequestContext) {
	namespaceIDStr := c.Param("id")
	userIDStr := c.Param("user_id")
	namespaceID, err := strconv.ParseInt(namespaceIDStr, 10, 64)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.ServiceErrCode,
			"message": err.Error(),
		})
		return
	}
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.ServiceErrCode,
			"message": err.Error(),
		})
		return
	}
	client, err := cloudfile.NewFileServiceClient()
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	err = client.RemoveNamespaceAuthority(ctx, userID, namespaceID)
	if err != nil {
		fmt.Println(err)
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.ServiceErrCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(consts.StatusOK, utils.H{
		"code":    errno.SuccessCode,
		"message": errno.SuccessMsg,
	})
}
