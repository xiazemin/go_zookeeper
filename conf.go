package main

import (
	"github.com/xiazemin/go_zookeeper/conf"
)

func main(){
	var  zkConf =new(conf.ZkConf)
	zkConf.Init(5)
	zkConf.Dispatch("confPath","[{a:0}]")
	zkConf.Dispatch("confPath1","[{a:1}]")
}