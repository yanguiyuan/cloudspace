// Code generated by Kitex v0.6.1. DO NOT EDIT.

package cloudfileservice

import (
	server "github.com/cloudwego/kitex/server"
	rpc "github.com/yanguiyuan/cloudspace/pkg/rpc"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler rpc.CloudFileService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
