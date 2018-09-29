package handle

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
)

var Hosts = []string{"127.0.0.1:2181","127.0.0.1:2182","127.0.0.1:2183"}//127.0.0.1:2888,127.0.0.1:2889,127.0.0.1:2890
var acls = zk.WorldACL(zk.PermAll)

const (
	PERSISTENT =0
	EPHEMERAL=1
	PERSISTENT_SEQUENTIAL=2
	EPHEMERAL_SEQUENTIAL=3

)

func Callback(event zk.Event) {
	fmt.Println("*******************")
	fmt.Println("path:", event.Path)
	fmt.Println("type:", event.Type.String())
	fmt.Println("state:", event.State.String())
	fmt.Println("-------------------")
}

func Create(conn *zk.Conn, path string, data []byte,flags int32 )string {
	val, err_create := conn.Create(path, data, flags, acls)
	if err_create != nil {
		fmt.Println(err_create)
		return val
	}
	fmt.Println("create:",val)
	return  val
}
