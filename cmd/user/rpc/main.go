package main

import (
	user "bbs/kitex_gen/user/userservice"
	"bbs/pkg/bound"
	"bbs/pkg/constants"
	"bbs/pkg/middleware"
	tracer2 "bbs/pkg/tracer"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"log"
	"net"
)

func Init() {
	tracer2.InitJaeger(constants.UserServiceName)
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	ip, port := constants.IP, constants.UserPort
	addr, err := net.ResolveTCPAddr("tcp", ip+port)
	if err != nil {
		panic(err)
	}
	Init()
	if err != nil {
		panic(err)
	}
	svr := user.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.UserServiceName}),
		server.WithMiddleware(middleware.CommonMiddleware),
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr), // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMuxTransport(), // Multiplex
		server.WithSuite(trace.NewDefaultServerSuite()),
		server.WithBoundHandler(bound.NewCpuLimitHandler()),
		server.WithRegistry(r),
	)

	if err = svr.Run(); err != nil {
		log.Println(err.Error())
	}
}
