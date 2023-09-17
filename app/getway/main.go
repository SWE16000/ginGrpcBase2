package main

import (
	"fmt"
	"ginGrpcBase2/app/getway/router"
	"ginGrpcBase2/config"
	"ginGrpcBase2/etcd"
	"ginGrpcBase2/pdgrpc"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/spf13/viper"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	config.InitConfig()
	startListen()
}

func startListen() {
	// 服务名
	userServiceName := viper.GetString("domain.user")
	fmt.Println("userServiceName",userServiceName)
	//SentinelTest()
	//TODO:etcd服务发现
	//连接etcd服务端
	client, err := etcd.NewEtcdClient("http://localhost:2379", userServiceName, 10)
	if err!=nil{
		log.Fatal("etcd服务端连接失败",err)
		panic("etcd服务端连接失败")
	}
	// RPC 连接
	connUser, err := RpcEtcdDail(client,userServiceName)
	if err != nil {
		return
	}
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	userServiceClient := pdgrpc.NewUserServiceClient(connUser)
	//TODO:加入熔断
	server:=router.SetRouter(userServiceClient)
	server.Run(":8081")
}
func RpcEtcdDail(c *etcd.EtcdClient,servername string)(conn *grpc.ClientConn, err error){
	etcdResolver, err := resolver.NewBuilder(c.Client)
	if err!=nil{
		log.Fatal("etcd服务发现失败",err)
		return nil,err
	}
	conn, err = grpc.Dial(fmt.Sprintf("etcd:///%s", servername),
		grpc.WithResolvers(etcdResolver),    // 注入 etcd resolver
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// 声明使用的负载均衡策略为 roundrobin
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))

	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	fmt.Println("addr: ", servername)
	fmt.Println("客户端连接错误：",err)
	if err != nil {
		log.Fatal("服务端出错，连接不上",err)
	}
	//退出时关闭连接
	//defer conn.Close()
	return
}

//grpc.Dial连接Grpc客户端
func RpcDial(addr string)(conn *grpc.ClientConn, err error){
	//conn, err = grpc.Dial(addr, grpc.WithInsecure())
	conn, err = grpc.Dial("127.0.0.1:7300", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	fmt.Println("addr: ", addr)
	fmt.Println("客户端连接错误：",err)
	if err != nil {
		log.Fatal("服务端出错，连接不上",err)
	}
	//退出时关闭连接
	//defer conn.Close()
	return
}

//func RPCConnect(ctx context.Context, serviceName string, etcdRegister *discovery.Resolver) (conn *grpc.ClientConn, err error) {
//	opts := []grpc.DialOption{
//		grpc.WithInsecure(),
//	}
//	addr := fmt.Sprintf("%s:///%s", etcdRegister.Scheme(), serviceName)
//	conn, err = grpc.DialContext(ctx, addr, opts...)
//	return
//}

func SentinelTest(){
	//err:=sentinel.InitDefault()
	err:=sentinel.InitWithConfigFile("/ginGrpcBase2/config/sentinel.yaml")

	if err != nil {
		log.Fatalf("Unexpected error: %+v", err)
	}
	// 配置一条限流规则
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "test",
			Threshold:              2,
			StatIntervalInMs:      1000,
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

}
