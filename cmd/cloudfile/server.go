package main

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/yanguiyuan/cloudspace/internal/cloudfile/dal"
	"github.com/yanguiyuan/cloudspace/internal/cloudfile/handler"
	rpc "github.com/yanguiyuan/cloudspace/pkg/rpc/cloudfileservice"
	"github.com/yanguiyuan/yuan/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	c := config.NewConfig()
	db, err := gorm.Open(mysql.Open(c.GetString("cloudfile.mysql.dsn")), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
	}
	dal.SetDefault(db)
	client, err := oss.New(c.GetString("cloudfile.oss.endpoint"), c.GetString("cloudfile.oss.accessKeyID"), c.GetString("oss.accessKeySecret"))
	if err != nil {
		log.Println(err.Error())
	}
	// 填写存储空间名称。
	bucket, err := client.Bucket(c.GetString("cloudfile.oss.bucketName"))
	service := handler.CloudFileServiceImpl{OssBucket: bucket}
	svr := rpc.NewServer(&service)
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
