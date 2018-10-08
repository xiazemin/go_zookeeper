package lock

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
	"github.com/satori/go.uuid"

	"github.com/xiazemin/go_zookeeper/handle"
)

type ZkLock struct{
	BasePath string
	Data []byte
	Flag int32
	Conn *zk.Conn
}
func (this*ZkLock)Init(second time.Duration){
	this.BasePath = "/lock"
	this.Data=[]byte("zk lock root")
	this.Flag=handle.PERSISTENT
	conn, connEvent, err := zk.Connect(handle.Hosts, time.Second*second)
	this.Conn=conn

	exist, _, ech, err := this.Conn.ExistsW(this.BasePath)
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
	defer this.Conn.Close()
	handle.Create(this.Conn, this.BasePath ,this.Data,this.Flag)
}
func watchCreataNode(ech <-chan zk.Event){
	event:=<-ech
	fmt.Println("*******watch************")
	fmt.Println("path:", event.Path)
	fmt.Println("type:", event.Type.String())
	fmt.Println("state:", event.State.String())
	fmt.Println("-------------------")
}

func (this*ZkLock)Lock(key string)(bool,string){
	u,_:= uuid.NewV4()
	path:=this.BasePath+"/"+fmt.Sprintf("%s",u)+key
	var curNode string
	for true {
		curNode = handle.Create(this.Conn, path, this.Data, handle.EPHEMERAL_SEQUENTIAL)
		children, _, err := handle.GetChildren(this.Conn, this.BasePath)
		if err != nil {
			fmt.Println(err)
			break;
		}
		fmt.Println(curNode, children)
		if this.BasePath + "/" + children[0] == curNode {
			return true,curNode
		}

		if this.watchPrevious(children,curNode){
			break;
		}
	}
        return  false,curNode
}

func  (this*ZkLock)watchPrevious(children []string,curNode string) bool {
	i:=0
	child:=""
	for i,child=range children{
		if(this.BasePath+"/"+child)==curNode {
			break;
		}
	}
	fmt.Println("previous node:%s",children[i-1])
	exist, _, ech, err := this.Conn.ExistsW(this.BasePath+"/"+children[i-1])
	if exist{
		return true
	}
	if err != nil {
		fmt.Println(err)
		return false
	}
	go this.wait4lock(ech)
	return  false
}
func  (this*ZkLock)wait4lock(ech <-chan zk.Event)  {
	event:=<-ech
	fmt.Println("*******watch previous************")
	fmt.Println("path:", event.Path)
	fmt.Println("type:", event.Type.String())
	fmt.Println("state:", event.State.String())
	fmt.Println("-------------------")
}

func (this*ZkLock)UnLock(key string)bool  {
	err:=handle.Delete(this.Conn,key)
	if err!=nil{
		fmt.Println(err)
		return false
	}
	return  true
}