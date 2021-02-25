package broker

import (
	"unitalk/config"

	"github.com/go-redis/redis/v8"
)

// REDIS client
var REDIS *redis.ClusterClient

// init redis
func init() {
	conn := config.Config["redis"].(map[string]interface{})
	var sentinel []string
	for _, v := range conn["sentinelAddrs"].([]interface{}) {
		sentinel = append(sentinel, v.(string))
	}
	REDIS = redis.NewFailoverClusterClient(&redis.FailoverOptions{
		MasterName:     conn["masterName"].(string),
		SentinelAddrs:  sentinel,
		RouteByLatency: true,
	})
}
