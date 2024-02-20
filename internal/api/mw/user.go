package mw

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	_ "github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client"
	"github.com/yanguiyuan/cloudspace/pkg/errno"
	"github.com/yanguiyuan/cloudspace/pkg/rpc/cloudfileservice"
	"github.com/yanguiyuan/yuan/pkg/config"
	"strconv"
)

func NewCloudFileServiceClient() (cloudfileservice.Client, error) {
	return cloudfileservice.NewClient("cloudfile", client.WithHostPorts(config.GetString("cloudfile.addr")))
}

func CheckReadPermission(ctx context.Context, c *app.RequestContext) {
	// 获取用户id
	identity, b := c.Get(IdentityKey)
	if !b {
		c.String(consts.StatusUnauthorized, "未登录")
		c.Abort()
		return
	}
	//获取服务
	client, err := NewCloudFileServiceClient()
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}
	//获取namespaceID
	fileID := c.Param("id")
	//获取权限等级
	authority, a := client.GetAuthority(ctx, identity.(int64), fileID)
	if a != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.ServiceErrCode,
			"message": a.Error(),
		})
		c.Abort()
		return
	}
	// 只有1，2,3有°权限
	if authority < 1 || authority > 3 {
		// 如果没有用户信息，返回未授权错误
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.NoReadPermissionCode,
			"message": errno.NoReadPermissionMsg,
		})
		// 终止请求处理流程
		c.Abort()
		return
	}
}

func CheckWritePermission(ctx context.Context, c *app.RequestContext) {
	// 获取用户id
	identity, b := c.Get(IdentityKey)
	if !b {
		c.String(consts.StatusUnauthorized, "未登录")
		c.Abort()
		return
	}
	//获取服务
	client, err := NewCloudFileServiceClient()
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}
	//获取fileID
	fileID := c.Param("id")
	//获取权限等级
	authority, a := client.GetAuthority(ctx, identity.(int64), fileID)
	if a != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.ServiceErrCode,
			"message": a.Error(),
		})
		c.Abort()
		return
	}
	// 只有1和2具备写权限
	if !(authority == 1 || authority == 2) {
		// 如果没有用户信息，返回未授权错误
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.NoWritePermissionCode,
			"message": errno.NoWritePermissionMsg,
		})
		// 终止请求处理流程
		c.Abort()
		return
	}
}

func CheckAdminPermission(ctx context.Context, c *app.RequestContext) {
	// 获取用户id
	identity, b := c.Get(IdentityKey)
	if !b {
		c.String(consts.StatusUnauthorized, "未登录")
		c.Abort()
		return
	}
	//获取用户服务
	client, err := NewUserServiceClient()
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}
	//获取用户信息
	resp, err := client.GetUser(ctx, identity.(int64))
	if err != nil {
		c.String(consts.StatusUnauthorized, err.Error())
	}
	//获取用户等级
	role := resp.Role
	//如果用户角色是user则不能访问
	if role == "user" {
		// 如果没有用户信息，返回未授权错误
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.NoAdminPermissionCode,
			"message": "No admin permission",
		})
		// 终止请求处理流程
		c.Abort()
		return
	}
}

func CheckAllPermission(ctx context.Context, c *app.RequestContext) {
	// 获取用户id
	identity, b := c.Get(IdentityKey)
	if !b {
		c.String(consts.StatusUnauthorized, "未登录")
		c.Abort()
		return
	}
	//获取服务
	client, err := NewCloudFileServiceClient()
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}
	//获取namespaceID
	namespaceID := c.Param("id")
	parseInt, err := strconv.ParseInt(namespaceID, 10, 64)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.InvalidParamCode,
			"message": "解析namespaceID失败",
		})
		c.Abort()
		return
	}
	//获取权限等级
	authority, a := client.QueryUserNamespaceAuthority(ctx, identity.(int64), parseInt)
	if a != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.ServiceErrCode,
			"message": a.Error(),
		})
		c.Abort()
		return
	}
	//如果权限不为0，就没有授权权限
	if authority != 1 {
		// 如果没有用户信息，返回未授权错误
		c.JSON(consts.StatusOK, utils.H{
			"code":    errno.NoAllPermissionCode,
			"message": errno.NoAllPermissionMsg,
		})
		// 终止请求处理流程
		c.Abort()
		return
	}
}
