package conf

import (
	"fmt"
	"time"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/xiazemin/go_zookeeper/handle"
)

type ZkConf struct{
	BasePath string
	Data []byte
	Flag int32
	Conn *zk.Conn
}
func (this*ZkConf)Init(second time.Duration){
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
	handle.Create(this.Conn, this.BasePath ,this.Data,this.Flag)
}

func (this *ZkConf)Dispatch(path string, data string)bool{
	exist,stat, err:=handle.Exist(this.Conn,this.BasePath+"/"+path)
	if err!=nil{
		fmt.Println(stat)
          return  false
	}
	if !exist{
		handle.Create(this.Conn,this.BasePath+"/"+path,[]byte(data),handle.PERSISTENT)
		return true
	}
	dataOld,stat,err:=handle.Get(this.Conn,this.BasePath+"/"+path)
	fmt.Println("dataOld:%+v",dataOld)
	handle.Update(this.Conn,this.BasePath+"/"+path,[]byte(data),stat.Version)
	return true
}

func watchCreataConfNode(ech <-chan zk.Event) {
	event := <-ech
	fmt.Println("*******watch************")
	fmt.Println("path:", event.Path)
	fmt.Println("type:", event.Type.String())
	fmt.Println("state:", event.State.String())
	fmt.Println("-------------------")
}
