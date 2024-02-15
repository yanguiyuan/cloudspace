package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/yanguiyuan/yuan/pkg/config"
	"log"
)

func main() {
	c := config.NewConfig()
	client, err := oss.New("https://oss-cn-hangzhou.aliyuncs.com", "LTAI5tCZenyLGTYY2NDGP2Ge", "eN32dnGnJ6XqeehDgkeXjwW4384cXA")
	if err != nil {
		log.Println("ee", err.Error())
	}
	b, err := client.Bucket(c.GetString("cloudfile.oss.bucketName"))
	if err != nil {
		log.Println("ee", err.Error())
	}
	res, err := b.ListObjects()
	fmt.Println("my objects:", getObjectsFormResponse(res))
	//marker := ""
	//for {
	//	lsRes, err := client.ListBuckets(oss.Marker(marker))
	//	if err != nil {
	//		fmt.Println("Error:", err)
	//		os.Exit(-1)
	//	}
	//
	//	// 默认情况下一次返回100条记录。
	//	for _, bucket := range lsRes.Buckets {
	//		fmt.Println("Bucket: ", bucket.Name)
	//	}
	//
	//	if lsRes.IsTruncated {
	//		marker = lsRes.NextMarker
	//	} else {
	//		break
	//	}
	//}
}
func getObjectsFormResponse(lor oss.ListObjectsResult) string {
	var output string
	for _, object := range lor.Objects {
		output += object.Key + "  "
	}
	return output
}
