package main
import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
	"github.com/xiazemin/go_zookeeper/handle"
)

var path1 = "/whatzk"
var data1 = []byte("hello,this is a zk go test demo!!!")
//var event chan zk.Event =make(chan zk.Event,1)
func main() {
	conn, _, err := zk.Connect(handle.Hosts, time.Second*5)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	_, _, ech, err := conn.ExistsW(path1)
	if err != nil {
		fmt.Println(err)
		return
	}

	handle.Create(conn, path1, data1,zk.FlagEphemeral)

	go watchCreataNode(ech)

}

func watchCreataNode(ech <-chan zk.Event){
	event:=<-ech
	fmt.Println("*******watch************")
	fmt.Println("path:", event.Path)
	fmt.Println("type:", event.Type.String())
	fmt.Println("state:", event.State.String())
	fmt.Println("-------------------")
}