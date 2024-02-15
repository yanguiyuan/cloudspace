package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/yanguiyuan/cloudspace/internal/api/mw"
)

func Ping(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, utils.H{
		"message": "pong",
	})
}
func ExtractID(c *app.RequestContext) (int64, bool) {
	identity, ok := c.Get(mw.IdentityKey)
	if !ok {
		c.JSON(consts.StatusBadRequest, utils.H{
			"message": "identity not found",
		})
		return 0, false
	}
	return identity.(int64), true
}
