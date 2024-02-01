package file

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/yanguiyuan/cloudspace/internal/api/mw"
	"github.com/yanguiyuan/cloudspace/internal/cloudfile"
	"github.com/yanguiyuan/cloudspace/pkg/errno"
	"log"
)

// Query .
// @router /files/:parent_id [GET]
func QueryFileItemList(ctx context.Context, c *app.RequestContext) {
	parentID := c.Query("parent_id")
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
	}
	c.JSON(consts.StatusOK, utils.H{
		"code":    errno.SuccessCode,
		"message": errno.SuccessMsg,
		"data":    r,
	})
}

// Upload .
// @router /files [POST]
func Upload(ctx context.Context, c *app.RequestContext) {
	form, _ := c.MultipartForm()
	files := form.File["data"]

	for _, file := range files {
		fmt.Println(file.Filename)
		// Upload the file to specific dst.
		err := c.SaveUploadedFile(file, fmt.Sprintf("./%s", file.Filename))
		if err != nil {
			log.Fatal(err)
		}
	}
	c.String(consts.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
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
