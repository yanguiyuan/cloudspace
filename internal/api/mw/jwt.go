package mw

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client"
	"github.com/hertz-contrib/jwt"
	"github.com/yanguiyuan/cloudspace/internal/api/model/user"
	"github.com/yanguiyuan/cloudspace/pkg/errno"
	"github.com/yanguiyuan/cloudspace/pkg/rpc/userservice"
	"github.com/yanguiyuan/yuan/pkg/config"
	"log"
	"net/http"
	"time"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	IdentityKey   = "identity"
)

func NewUserServiceClient() (userservice.Client, error) {
	return userservice.NewClient("user", client.WithHostPorts(config.GetString("user.addr")))
}
func InitJWT() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:         "test zone",
		Key:           []byte(config.GetString("api.jwt.secret")),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		IdentityKey:   IdentityKey,
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, utils.H{
				"code":  errno.SuccessCode,
				"msg":   errno.SuccessMsg,
				"token": token,
			})
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return int64(claims[IdentityKey].(float64))

		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var req user.RegisterReq
			if err := c.BindAndValidate(&req); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			client, err := NewUserServiceClient()
			if err != nil {
				c.String(consts.StatusInternalServerError, err.Error())
			}
			r, err := client.UserLogin(ctx, req.Username, req.Password)
			if err == nil {
				return r, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
}
