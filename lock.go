package main

import (
	"github.com/xiazemin/go_zookeeper/lock"

	"fmt"
)
func main() {
	var zkl = new(lock.ZkLock)
	zkl.Init(5)
	// 创建

	res1,key1:=zkl.Lock("lockTest")
	res2,key2:=zkl.Lock("lockTest1")
	fmt.Println(res1,res2)
	zkl.UnLock(key2)
	zkl.UnLock(key1)
}
