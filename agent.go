package main

import (
	"fmt"
	"github.com/xiazemin/go_zookeeper/conf"
	"time"
)

func main() {
	var zkAgent = new(conf.ZkWatch)
	//会话时间（Session Time）
	//在《ZooKeeper API 使用》一文中已经提到，在实例化一个ZK客户端的时候，需要设置一个会话的超时时间。这里需要注意的一点是，客户端并不是可以随意设置这个会话超时时间，在ZK服务器端对会话超时时间是有限制的，主要是minSessionTimeout和maxSessionTimeout这两个参数设置的。（详细查看这个文章《ZooKeeper管理员指南》）Session超时时间限制，如果客户端设置的超时时间不在这个范围，那么会被强制设置为最大或最小时间。
	//默认的Session超时时间是在2 * tickTime ~ 20 * tickTime。所以，如果应用对于这个会话超时时间有特殊的需求的话，一定要和ZK管理员沟通好，确认好服务端是否设置了对会话时间的限制。
	for ; ;  {
		zkAgent.Init(20)
		data, version := zkAgent.Watch("confPath")
		fmt.Println(string(data))
		fmt.Println(version)
		time.Sleep(time.Second*20)
	}
}
