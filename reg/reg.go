package reg

import (
	"time"
	"unitalk/config"

	"github.com/go-zookeeper/zk"
)

// ZK connection
var ZK *zk.Conn

func init() {
	zkConf := config.Config["zk"].(map[string]interface{})
	var zkServers []string
	for _, v := range zkConf["serverList"].([]interface{}) {
		zkServers = append(zkServers, v.(string))
	}
	ZK, _, err := zk.Connect(zkServers, 15*time.Second)
	if err != nil {
		panic(err)
	}
	ZK.Create(zkConf["path"].(string)+config.Config["listen"].(string), nil,
		zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
}
