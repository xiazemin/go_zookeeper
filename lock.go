package main

import (
	"github.com/xiazemin/go_zookeeper/lock"
)
func main() {
	var zkl = new(lock.ZkLock)
	zkl.Init()
}

