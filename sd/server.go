package sd
//这个结构数据会存储在节点的 data 中，表示服务发现的地址信息。
type ServiceNode struct {
	Name string `json:"name"` // 服务名称，这里是 user
	Host string `json:"host"`
	Port int    `json:"port"`
}
