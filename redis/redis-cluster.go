package redis

import (
    "log"
    "github.com/go-redis/redis"
)

var client * redis.ClusterClient

func InitRedis(_addrs []string, _prefix string)  {
    re := redis.NewClusterClient(&redis.ClusterOptions{
        Addrs:_addrs,
        PoolSize:1000,
    })
    if err := re.ReloadState();err != nil {
        log.Fatal("reload redis cluster state failed")
    }
    client = re
}

func GetClient() *redis.ClusterClient {
    return client
}