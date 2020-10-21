package redis

import (
    "fmt"
    "github.com/go-redis/redis"
    "log"
)

var cli *redis.Client

func InitRedisStand(_addr string,_poolSize,_maxIdleConns int)  {
    standClient := redis.NewClient(&redis.Options{
        Addr:_addr,
        PoolSize:_poolSize,
        MinIdleConns:_maxIdleConns,
    })
    pong,err := standClient.Ping().Result()
    if err != nil {
        log.Fatal(fmt.Sprintf("Init Redis Failed,err:%v",err))
    }
    log.Printf("Init Redis Stand Client success!,pong:%s",pong)
}

func GetRedisStandClient() *redis.Client {
    return cli
}