package file

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/yanguiyuan/cloudspace/internal/api/handler"
	"github.com/yanguiyuan/cloudspace/internal/api/model/file"
	"github.com/yanguiyuan/cloudspace/internal/api/mw"
	"github.com/yanguiyuan/cloudspace/internal/cloudfile"
	"github.com/yanguiyuan/cloudspace/pkg/errno"
	"github.com/yanguiyuan/cloudspace/pkg/rpc"
	"io"
	"strconv"
)

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
	identity, b := c.Get(mw.IdentityKey)
	if !b {
		c.String(consts.StatusUnauthorized, "未登录")
	}
	fmt.Println("pd:", parentID)
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
		item, err = client.Add(ctx, &rpc.AddRequest{
			FileName: file.Filename,
			ParentID: parentID,
			Uid:      identity.(int64),
			FileData: d,
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

// UserRootFile .
// @router /user/file/root [GET]
func UserRootFile(ctx context.Context, c *app.RequestContext) {
	if identity, ok := c.Get(mw.IdentityKey); ok {
		client, err := cloudfile.NewFileServiceClient()
		if err != nil {
			c.String(consts.StatusInternalServerError, err.Error())
			return
		}
		r, err := client.QueryUserFileRoot(ctx, identity.(int64))
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
			"id":      r,
		})
		return
	}
	c.String(consts.StatusUnauthorized, "未登录")
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
	if identity, ok := c.Get(mw.IdentityKey); ok {
		client, err := cloudfile.NewFileServiceClient()
		if err != nil {
			c.String(consts.StatusInternalServerError, err.Error())
			return
		}
		resp, err = client.GetFileURL(ctx, id, identity.(int64))
		if err != nil {
			c.JSON(consts.StatusOK, utils.H{
				"code":    errno.NotFoundCode,
				"message": errno.NotFoundMsg,
			})
			return
		}
	}
	c.JSON(consts.StatusOK, utils.H{
		"code":    errno.SuccessCode,
		"message": errno.SuccessMsg,
		"data":    resp,
	})
}
func DeleteFileItem(ctx context.Context, c *app.RequestContext) {
	var req file.DeleteFileItemReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.InvalidParamCode,
			"message": errno.InvalidParamMsg,
		})
		return
	}
	identity, b := c.Get(mw.IdentityKey)
	if !b {
		c.String(consts.StatusUnauthorized, "未登录")
	}
	client, err := cloudfile.NewFileServiceClient()
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	err = client.Remove(ctx, &rpc.RemoveRequest{
		Id:       req.ID,
		Uid:      identity.(int64),
		Filename: req.FileName,
	})
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
	identity, b := c.Get(mw.IdentityKey)
	if !b {
		c.String(consts.StatusUnauthorized, "未登录")
	}
	fmt.Println("userID:", identity)

	client, err := cloudfile.NewFileServiceClient()
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	resp, err := client.QueryUserNamespaces(ctx, identity.(int64))
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
	r, err := client.CreateFileItem(ctx, name, "namespace", "@root")
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.ServiceErrCode,
			"message": err.Error(),
		})
		return
	}
	namespaceID, err := client.CreateNamespace(ctx, name, r)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.ServiceErrCode,
			"message": err.Error(),
		})
		return
	}
	err = client.CreateUserNamespace(ctx, userid, namespaceID, 0)
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

func LinkNamespace(ctx context.Context, c *app.RequestContext) {
	c.Redirect(consts.StatusOK, []byte("http://localhost:5173/link?user_id="))
	id, b := handler.ExtractID(c)
	namespaceId := c.Param("id")
	if !b {
		return
	}
	token, _, err := mw.JwtMiddleware.TokenGenerator(utils.H{
		"namespace_id": namespaceId,
		"uid":          id,
	})
	if err != nil {
		return
	}
	c.JSON(consts.StatusOK, utils.H{
		"code":    errno.SuccessCode,
		"message": errno.SuccessMsg,
		"data": "http://localhost:5173/link?user_id=" +
			strconv.FormatInt(id, 10) +
			"&namespace_id=" + namespaceId +
			"&token=" + token,
	})
}
