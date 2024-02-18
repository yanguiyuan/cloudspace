package main

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cloudwego/kitex/server"
	"github.com/yanguiyuan/cloudspace/internal/cloudfile/dal"
	"github.com/yanguiyuan/cloudspace/internal/cloudfile/handler"
	rpc "github.com/yanguiyuan/cloudspace/pkg/rpc/cloudfileservice"
	"github.com/yanguiyuan/yuan/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net"
)

func main() {
	c := config.NewConfig()
	db, err := gorm.Open(mysql.Open(c.GetString("cloudfile.mysql.dsn")), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
	}
	dal.SetDefault(db)
	client, err := oss.New(c.GetString("cloudfile.oss.endpoint"), c.GetString("cloudfile.oss.accessKeyID"), c.GetString("cloudfile.oss.accessKeySecret"))
	if err != nil {
		log.Println(err.Error())
	}
	// 填写存储空间名称。
	bucket, err := client.Bucket(c.GetString("cloudfile.oss.bucketName"))
	service := handler.CloudFileServiceImpl{OssBucket: bucket}
	addr, _ := net.ResolveTCPAddr("tcp", c.GetString("cloudfile.addr"))
	svr := rpc.NewServer(&service, server.WithServiceAddr(addr), server.WithReadWriteTimeout(1000*1000*60*50))
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
