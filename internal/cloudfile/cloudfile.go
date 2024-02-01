package cloudfile

import (
	"github.com/cloudwego/kitex/client"
	"github.com/yanguiyuan/cloudspace/pkg/rpc/cloudfileservice"
	"github.com/yanguiyuan/yuan/pkg/config"
)

func NewFileServiceClient() (cloudfileservice.Client, error) {
	return cloudfileservice.NewClient("cloudfile", client.WithHostPorts(config.GetString("cloudfile.addr")))
}
