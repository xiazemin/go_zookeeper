package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/xiazemin/go_zookeeper/handle"
)

func main(){
	//使用步骤如下：（相关代码位于dnshostprovider.go中）
	hostPro:=new(zk.DNSHostProvider)
	err:=hostPro.Init(handle.Hosts)//先初始化
	if err != nil {
		fmt.Println(err)
		return
	}
	server,retryStart:=hostPro.Next()//获得host
	fmt.Println(server,retryStart)

	hostPro.Connected()  //连接成功后会调用
	//上面的一系列步骤都集成在func Connect(servers []string, sessionTimeout time.Duration, options ...connOption) (*Conn, <-chan Event, error)中
}