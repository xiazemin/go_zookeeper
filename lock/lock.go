package lock

import (
	"fmt"
	"time"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/xiazemin/go_zookeeper/handle"
)

type ZkLock struct{
	BasePath string
	Data []byte
	Flag int32
}
func (this*ZkLock)Init(){
	this.BasePath = "/lock"
	this.Data=[]byte("zk lock root")
	this.Flag=handle.PERSISTENT
	conn, connEvent, err := zk.Connect(handle.Hosts, time.Second*5)

	exist, _, ech, err := conn.ExistsW(this.BasePath)
	if exist{
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}

	go watchCreataNode(ech)

	if err != nil {
		fmt.Println(err,connEvent)
		return
	}
	defer conn.Close()
	handle.Create(conn, this.BasePath ,this.Data,this.Flag)
}
func watchCreataNode(ech <-chan zk.Event){
	event:=<-ech
	fmt.Println("*******watch************")
	fmt.Println("path:", event.Path)
	fmt.Println("type:", event.Type.String())
	fmt.Println("state:", event.State.String())
	fmt.Println("-------------------")
}