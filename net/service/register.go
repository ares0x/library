package service

import (
    "github.com/coreos/etcd/clientv3"
    "log"
    "time"
)

type ServiceInfo struct {
    ID  uint64
    IP  string
}

type Service struct {
    Name string
    Info ServiceInfo
    stop chan error
    leaseid clientv3.LeaseID
    client *clientv3.Client
}

func NewService(name string, info ServiceInfo, endpoints []string) (*Service, error) {
    cli,err := clientv3.New(clientv3.Config{
        Endpoints:endpoints,
        DialTimeout: 2 * time.Second,
    })
    if err != nil {
        log.Fatalf("init error:%v",err)
        return nil,err
    }
    return &Service{
        Name:name,
        Info:info,
        stop:make(chan error),
        client:cli,
    },err
}