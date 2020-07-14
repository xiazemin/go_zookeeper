package main
import (
	"fmt"
	"github.com/xiazemin/go_zookeeper/election"
)

func main() {
	// zookeeper配置
	zkConfig := &election.ZookeeperConfig{
		Servers:    []string{"127.0.0.1:2181"},
		RootPath:   "/ElectMasterDemo",
		MasterPath: "/master",
	}
	// main goroutine 和 选举goroutine之间通信的channel，同于返回选角结果
	isMasterChan := make(chan bool)

	var isMaster bool

	// 选举
	electionManager :=election. NewElectionManager(zkConfig, isMasterChan)
	go electionManager.Run()

	for {
		select {
		case isMaster = <-isMasterChan:
			if isMaster {
				// do some job on master
				fmt.Println("do some job on master")
			}
		}
	}
}