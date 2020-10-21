package rpc

import (
    "google.golang.org/grpc"
    "google.golang.org/grpc/keepalive"
    "log"
    "net"
    "sync"
    "time"
)

var (
    _defaultSerConf = &ServerConfig{
        Network:           "tcp",
        Addr:              "0.0.0.0:9000",
        Timeout:           time.Duration(time.Second),
        IdleTimeout:       time.Duration(time.Second * 60),
        MaxLifeTime:       time.Duration(time.Hour * 2),
        ForceCloseWait:    time.Duration(time.Second * 20),
        KeepAliveInterval: time.Duration(time.Second * 60),
        KeepAliveTimeout:  time.Duration(time.Second * 20),
    }
)

type ServerConfig struct {
    Network string `dsn:"network"`
    Addr string `dsn:"address"`
    Timeout time.Duration `dsn:"query.timeout"`
    IdleTimeout time.Duration `dsn:"query.idleTimeout"`
    MaxLifeTime time.Duration `dsn:"query.maxLife"`
    ForceCloseWait time.Duration `dsn:"query.closeWait"`
    KeepAliveInterval time.Duration `dsn:"query.keepaliveInterval"`
    KeepAliveTimeout time.Duration `dsn:"query.keepaliveTimeout"`
}

type Server struct {
    conf  *ServerConfig
    mutex sync.RWMutex
    server   *grpc.Server
    handlers []grpc.UnaryServerInterceptor
}

func (s *Server) SetDefaultConf(conf *ServerConfig) (error) {
    if conf == nil{
        conf = _defaultSerConf
    }
    if conf.Timeout <= 0 {
        conf.Timeout = time.Duration(time.Second)
    }
    if conf.IdleTimeout <= 0 {
        conf.IdleTimeout = time.Duration(time.Second * 60)
    }
    if conf.MaxLifeTime <= 0 {
        conf.MaxLifeTime = time.Duration(time.Hour * 2)
    }
    if conf.ForceCloseWait <= 0 {
        conf.ForceCloseWait = time.Duration(time.Second * 20)
    }
    if conf.KeepAliveInterval <= 0 {
        conf.KeepAliveInterval = time.Duration(time.Second * 60)
    }
    if conf.KeepAliveTimeout <= 0 {
        conf.KeepAliveTimeout = time.Duration(time.Second * 20)
    }
    if conf.Addr == "" {
        conf.Addr = "0.0.0.0:9000"
    }
    if conf.Network == "" {
        conf.Network = "tcp"
    }
    s.mutex.Lock()
    s.conf = conf
    s.mutex.Unlock()
    return nil
}

func NewServer(conf *ServerConfig) (s *Server) {
    s = new(Server)
    if err := s.SetDefaultConf(conf);err != nil{
        log.Fatal("load config failed")
    }
    options := []grpc.ServerOption{}
    keepParam := grpc.KeepaliveParams(keepalive.ServerParameters{
        MaxConnectionIdle:     time.Duration(s.conf.IdleTimeout),
        MaxConnectionAgeGrace: time.Duration(s.conf.ForceCloseWait),
        Time:                  time.Duration(s.conf.KeepAliveInterval),
        Timeout:               time.Duration(s.conf.KeepAliveTimeout),
        MaxConnectionAge:      time.Duration(s.conf.MaxLifeTime),
    })
    options = append(options,keepParam)
    s.server = grpc.NewServer(options...)
    return
}

func (s *Server) Server() *grpc.Server{
    return s.server
}

func (s *Server) Serve(lis net.Listener) error {
    return s.server.Serve(lis)
}

func (s *Server) Start() (*Server, error) {
    listen,err := net.Listen(s.conf.Network, s.conf.Addr)
    if err != nil {
        return nil,err
    }
    go func() {
        if err := s.Serve(listen);err != nil {
            panic(err)
        }
    }()
    return s,nil
}