package sd

import (
	"encoding/json"
	"github.com/samuel/go-zookeeper/zk")
//实现消费者获取服务列表方法
//获取服务节点列表时，我们先获取字节点的名称列表，然后依次读取内容拿到服务地址。因为获取字节点名称和获取字节点内容不是一个原子操作，所以在调用 Get 获取内容时可能会出现节点不存在错误，这是正常现象。
func (s *SdClient) GetNodes(name string) ([]*ServiceNode, error) {
	path := s.zkRoot + "/" + name
	// 获取字节点名称
	childs, _, err := s.conn.Children(path)
	if err != nil {
		if err == zk.ErrNoNode {
			return []*ServiceNode{}, nil
		}
		return nil, err
	}
	nodes := []*ServiceNode{}
	for _, child := range childs {
		fullPath := path + "/" + child
		data, _, err := s.conn.Get(fullPath)
		if err != nil {
			if err == zk.ErrNoNode {
				continue
			}
			return nil, err
		}
		node := new(ServiceNode)
		err = json.Unmarshal(data, node)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}
