package etcd

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	gresolver "google.golang.org/grpc/resolver"
	"log"
)

type EtcdClient struct{
	Client *clientv3.Client
	url string
	servername string
	leasettl int64   //租约有效期
	leaseid  clientv3.LeaseID //租约id
}

func NewEtcdClient(url string,servername string,leasettl int64)(* EtcdClient,error){
	etcdClient, err := clientv3.NewFromURL(url)
	if err!=nil{
		log.Fatal("创建etcd客户端失败",err)
		return nil,err
	}
	return &EtcdClient{
		Client:etcdClient,
		url:url,
		servername: servername,
		leasettl: leasettl,
	},nil
}

func (c*EtcdClient)Register(addr string )error{
	em, err := endpoints.NewManager(c.Client, c.servername)
	if err != nil {
		return err
	}

	lease, _ := c.Client.Grant(context.TODO(), c.leasettl)
	c.leaseid=lease.ID
	err = em.AddEndpoint(context.TODO(),
		fmt.Sprintf("%s/%s", c.servername, addr),
		endpoints.Endpoint{Addr: addr},
		clientv3.WithLease(lease.ID))
	if err != nil {
		return err
	}
	return nil
}

func (c*EtcdClient)KeepAlive()error{
	//etcdClient.KeepAlive(context.TODO(), lease.ID)
	alive, err := c.Client.KeepAlive(context.TODO(), c.leaseid)
	if err != nil {
		return err
	}

	go func() {
		for {
			<-alive
			fmt.Println("etcd server keep alive")
		}
	}()
	return nil
}
//删除节点
func (c*EtcdClient)UnRegister(addr string) error {
	log.Printf("etcdUnRegister %s\b", addr)
	if c.Client != nil {
		em, err := endpoints.NewManager(c.Client, c.servername)
		if err != nil {
			return err
		}
		err = em.DeleteEndpoint(context.TODO(), fmt.Sprintf("%s/%s", c.servername, addr))
		if err != nil {
			return err
		}
		return err
	}

	return nil
}
func (c*EtcdClient)Discover()(*gresolver.Builder,error){
	etcResolver, err := resolver.NewBuilder(c.Client)
	if err!=nil{
		return nil,err
	}
	return &etcResolver,nil
}