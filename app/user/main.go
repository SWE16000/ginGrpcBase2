package main

import (
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
	"ginGrpcBase2/config"
	"ginGrpcBase2/etcd"
	"ginGrpcBase2/app/user/service"
	"ginGrpcBase2/pdgrpc"
)

func main()  {
	config.InitConfig()
	server:=grpc.NewServer()
	//defer server.Stop()
	grpcAddress:=viper.GetString("server.grpcAddress")
	fmt.Println("grpcAddress",grpcAddress)
	//连接etcd服务端
	client, err := etcd.NewEtcdClient("http://localhost:2379", "user", 10)
	if err!=nil{
		log.Fatal("etcd服务端连接失败",err)
		panic("etcd服务端连接失败")
	}
	//TODO:etcd服务注册
	client.Register(grpcAddress)
	client.KeepAlive()
	listen, err := net.Listen("tcp", grpcAddress)
	if err!=nil{
		log.Fatal("服务监听失败")
		panic("服务监听失败")
	}
	//注册定义服务主体
	pdgrpc.RegisterUserServiceServer(server, service.UserService{})
	err = server.Serve(listen)
	if err!=nil{
		log.Fatal("服务启动失败")
	}
	fmt.Println("--------服务启动成功----------")
}
