package conf

import (
	"fmt"
	"time"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/xiazemin/go_zookeeper/handle"
)

type ZkWatch struct{
	BasePath string
	Data []byte
	Flag int32
	Conn *zk.Conn
}
func (this*ZkWatch)Init(second time.Duration){
	this.BasePath = "/conf"
	this.Data=[]byte("zk config dispatch root")
	this.Flag=handle.PERSISTENT
	conn, connEvent, err := zk.Connect(handle.Hosts, time.Second*second)
	this.Conn=conn

	exist, stat, ech, err := this.Conn.ExistsW(this.BasePath)
	if exist{
		fmt.Println(stat)
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}

	go watchCreataConfNode(ech)

	if err != nil {
		fmt.Println(err,connEvent)
		return
	}
	defer this.Conn.Close()
}

func (this*ZkWatch)Watch(path string)(data []byte,version int32)  {
	exist, stat, ech, err := this.Conn.ExistsW(this.BasePath+"/"+path)
	if err!=nil{
		fmt.Println("exist watch error:%+v",ech)
		return nil,-1
	}
	if !exist{
		fmt.Println("not exist:%+v",stat)
		go this.exitsEventWatch(ech)
		return nil,-1
	}

	data, statG, echG, errG := this.Conn.GetW(this.BasePath+"/"+path)
	if errG!=nil{
		fmt.Println("get watch error:%+v",echG)
		return nil,-1
	}
	go this.updateEventWatch(echG)
	return data,statG.Version
}

func (this*ZkWatch)exitsEventWatch(ech <-chan zk.Event ){
	event:=<-ech
	fmt.Println("=========exist event watch ==========:\n")
	fmt.Println("path:", event.Path)
	fmt.Println("type:", event.Type.String())
	fmt.Println("state:", event.State.String())
	fmt.Println("-------------------")
	data,version:=this.Watch("confPath")
	fmt.Println(string(data))
	fmt.Println(version)

}


func (this*ZkWatch)updateEventWatch(ech <-chan zk.Event){
	event:=<-ech
	fmt.Println("========update event watch==========:\n")
	fmt.Println("path:", event.Path)
	fmt.Println("type:", event.Type.String())
	fmt.Println("state:", event.State.String())
	fmt.Println("-------------------")
	data,version:=this.Watch("confPath")
	fmt.Println(string(data))
	fmt.Println(version)
}