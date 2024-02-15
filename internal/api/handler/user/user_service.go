package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client"
	"github.com/yanguiyuan/cloudspace/internal/api/handler"
	"github.com/yanguiyuan/cloudspace/internal/api/model/user"
	"github.com/yanguiyuan/cloudspace/pkg/errno"
	"github.com/yanguiyuan/cloudspace/pkg/rpc"
	"github.com/yanguiyuan/cloudspace/pkg/rpc/userservice"
	"github.com/yanguiyuan/yuan/pkg/config"
	"strings"
)

func NewUserServiceClient() (userservice.Client, error) {
	return userservice.NewClient("user", client.WithHostPorts(config.GetString("user.addr")))
}

// Register .
// @router /register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.RegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	client, err := NewUserServiceClient()
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		hlog.Error(err)
		return
	}
	r, err := client.UserRegister(ctx, req.Username, req.Password)
	if err != nil {
		var respErr errno.ErrNo
		if strings.Contains(err.Error(), "Error 1062 (23000)") {
			respErr = errno.UserAlreadyExist
		} else {
			respErr = errno.ConvertErr(err)
			hlog.Error(err)
		}
		c.JSON(consts.StatusOK, utils.H{
			"code":    respErr.Code,
			"message": respErr.Msg,
		})
		return
	}
	c.JSON(consts.StatusOK, utils.H{
		"code":    errno.SuccessCode,
		"message": errno.SuccessMsg,
		"id":      r,
	})
}
func UserInfo(ctx context.Context, c *app.RequestContext) {
	userID, b := handler.ExtractID(c)
	if !b {
		return
	}
	client, err := NewUserServiceClient()
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		hlog.Error(err)
		return
	}
	r, err := client.GetUser(ctx, userID)
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
		"data":    r,
	})
}

func ModifyUserInfo(ctx context.Context, c *app.RequestContext) {
	var user rpc.User
	err := c.BindAndValidate(&user)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	userID, b := handler.ExtractID(c)
	if !b {
		return
	}
	if userID != user.Id {
		return
	}
	client, err := NewUserServiceClient()
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		hlog.Error(err)
		return
	}
	err = client.UpdateUser(ctx, &user)
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

func GetUsers(ctx context.Context, c *app.RequestContext) {
	var req user.GetUsersReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	client, err := NewUserServiceClient()
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		hlog.Error(err)
		return
	}
	r, err := client.GetUsers(ctx, req.Offset, req.Limit)
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
		"data":    r,
	})
}
