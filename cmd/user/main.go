package main

import (
	"github.com/cloudwego/kitex/server"
	"github.com/yanguiyuan/cloudspace/internal/user/dal"
	user "github.com/yanguiyuan/cloudspace/internal/user/handler"
	rpc "github.com/yanguiyuan/cloudspace/pkg/rpc/userservice"
	"github.com/yanguiyuan/yuan/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net"
)

func main() {
	c := config.NewConfig()
	db, err := gorm.Open(mysql.Open(c.GetString("user.mysql.dsn")), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
	}
	dal.SetDefault(db)
	addr, _ := net.ResolveTCPAddr("tcp", c.GetString("user.addr"))
	svr := rpc.NewServer(new(user.UserServiceImpl), server.WithServiceAddr(addr), server.WithReadWriteTimeout(1000*1000*10))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
