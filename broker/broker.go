package broker

import (
	"unitalk/config"

	"github.com/go-redis/redis/v8"
)

// REDIS client
var REDIS *redis.ClusterClient

// init redis
func init() {
	redisConf := config.Config["redis"].(map[string]interface{})
	var sentinel []string
	for _, v := range redisConf["sentinelAddrs"].([]interface{}) {
		sentinel = append(sentinel, v.(string))
	}
	REDIS = redis.NewFailoverClusterClient(&redis.FailoverOptions{
		MasterName:     redisConf["masterName"].(string),
		SentinelAddrs:  sentinel,
		RouteByLatency: true,
	})
}
