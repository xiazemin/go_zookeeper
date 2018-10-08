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

func GetChildren(conn *zk.Conn, path string)([]string, *zk.Stat, error)   {
	fmt.Println("getChilren:")
	return conn.Children(path)
}

func Delete(conn *zk.Conn, path string) error {
     fmt.Println("Delete:%s",path)
	return conn.Delete(path,0)
}

func  Get(conn *zk.Conn,path string)([]byte, *zk.Stat, error){
	fmt.Println("get data :%s\n",path)
	return  conn.Get(path)
}

func Exist(conn *zk.Conn,path string)(bool, *zk.Stat, error)  {
	fmt.Println("node exist:%s",path)
	return conn.Exists(path)
}

func Update(conn *zk.Conn,path string,data []byte,version  int32)(*zk.Stat, error)  {
	fmt.Println("update node:%s\t%s",path,string(data))
	return conn.Set(path,data,version)
}