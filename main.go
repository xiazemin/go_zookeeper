package main

import (
	"fmt"
	"time"
	"github.com/samuel/go-zookeeper/zk"

	"github.com/xiazemin/go_zookeeper/handle"
)

var path1 = "/go_zk"
var data1 = []byte("hello,this is a zk go test demo!!!")
func main() {
	option := zk.WithEventCallback(handle.Callback)

	conn, _, err := zk.Connect(handle.Hosts, time.Second*15, option)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	_, _, _, err = conn.ExistsW(path1)
	if err != nil {
		fmt.Println(err)
		return
	}

	handle.Create(conn, path1, data1,zk.FlagEphemeral)
	fmt.Println("get result:")
	conn.ExistsW(path1)
	fmt.Println()

	time.Sleep(time.Second * 2)

	_, _, _, err = conn.ExistsW(path1)
	if err != nil {
		fmt.Println(err)
		return
	}
	err=conn.Delete(path1, 0)
	fmt.Println(err)


}