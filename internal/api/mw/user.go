package mw

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	_ "github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client"
	"github.com/yanguiyuan/cloudspace/internal/cloudfile"
	"github.com/yanguiyuan/cloudspace/pkg/rpc/cloudfileservice"
	"github.com/yanguiyuan/yuan/pkg/config"
	"net/http"
	"strconv"
)

func NewCloudFileServiceClient() (cloudfileservice.Client, error) {
	return cloudfileservice.NewClient("cloudfile", client.WithHostPorts(config.GetString("cloudfile.addr")))
}

func CheckReadPermission(ctx context.Context, c *app.RequestContext) {
	userAuthority := c.Param("authority")
	identity, b := c.Get(IdentityKey) //*
	if !b {
		c.String(consts.StatusUnauthorized, "未登录")
	}
	client, err := cloudfile.NewFileServiceClient()
	resp, err := client.QueryUserNamespaces(ctx, identity.(int64)) //*
	if userAuthorityInt, err := strconv.Atoi(userAuthority); err == nil {
		if userAuthorityInt == 0 {
			fmt.Println("User is authorized to read the file")
			// 如果没有用户信息，返回未授权错误
			c.String(http.StatusUnauthorized, "Unauthorized") // 返回未授权错误
			// 终止请求处理流程
			c.Abort()
			return
		}
	}
}

func CheckWritePermission(ctx context.Context, c *app.RequestContext) {
	userAuthority := c.Param("authority")
	if userAuthorityInt, err := strconv.Atoi(userAuthority); err == nil {
		if userAuthorityInt == 0 {
			fmt.Println("User is authorized to read the file")
			// 如果没有用户信息，返回未授权错误
			c.String(http.StatusUnauthorized, "Unauthorized") // 返回未授权错误
			// 终止请求处理流程
			c.Abort()
			return
		}
	}
}
