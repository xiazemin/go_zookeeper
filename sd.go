package main

import (
	"fmt"

	"github.com/xiazemin/go_zookeeper/sd"
)
func main() {
	// 服务器地址列表
	servers := []string{"127.0.0.1:2181", "127.0.0.1:2182", "127.0.0.1:2183"}
	client, err := sd.NewClient(servers, "/sd", 10)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	node1 := &sd.ServiceNode{"user", "127.0.0.1", 4000}
	node2 := &sd.ServiceNode{"user", "127.0.0.1", 4001}
	node3 := &sd.ServiceNode{"user", "127.0.0.1", 4002}
	if err := client.Register(node1); err != nil {
		panic(err)
	}
	if err := client.Register(node2); err != nil {
		panic(err)
	}
	if err := client.Register(node3); err != nil {
		panic(err)
	}
	nodes, err := client.GetNodes("user")
	if err != nil {
		panic(err)
	}
	for _, node := range nodes {
		fmt.Println(node.Name,node.Host, node.Port)
	}
}