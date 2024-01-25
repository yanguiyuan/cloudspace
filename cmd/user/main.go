package main

import (
	user "github.com/yanguiyuan/cloudspace/internal/user/handler"
	rpc "github.com/yanguiyuan/cloudspace/pkg/rpc/userservice"
	"log"
)

func main() {
	svr := rpc.NewServer(new(user.UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
