package conf

import (
    "github.com/spf13/viper"
    _ "github.com/spf13/viper/remote"
    "log"
    "time"
)

var Vp *viper.Viper

func Init(provider,remoteAddr,confPath,confType string)  {
    v := viper.New()
    if err := v.AddRemoteProvider(provider, remoteAddr, confPath);err != nil {
        log.Printf("connect remote config repo failed,err:%v",err)
        panic("config")
    }
    v.SetConfigType(confType)
    if err := v.ReadRemoteConfig();err !=nil {
        log.Printf("get remote config failed,err:%v",err)
        return
    }
    Vp = v
    go func(vvv *viper.Viper) {
        for  {
            log.Printf("watche + 1")
            time.Sleep(time.Second * 5)
            err := vvv.WatchRemoteConfig()
            if err != nil {
                log.Printf("watch remote config failed,err:%v",err)
                continue
            }
        }
    }(v)
}

func GetVp() *viper.Viper {
    return Vp
}