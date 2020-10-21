package conf

import (
    "github.com/fsnotify/fsnotify"
    "github.com/spf13/viper"
)

type Setting struct {
    vp *viper.Viper
}

func NewSetting(configName string,configs ...string) (*Setting, error) {
    vp := viper.New()
    vp.SetConfigName(configName)
    for _, config := range configs {
        if config != "" {
            vp.AddConfigPath(config)
        }
    }
    vp.SetConfigType("yaml")
    err := vp.ReadInConfig()
    if err != nil {
        return nil, err
    }
    s := &Setting{vp: vp}
    s.WatchSettingConfigChange()
    return s,nil
}

func (s *Setting) WatchSettingConfigChange() {
    go func() {
        s.vp.WatchConfig()
        s.vp.OnConfigChange(func(in fsnotify.Event) {
            _ =s.ReloadAllSection()
        })
    }()
}

var sections = make(map[string]interface{})

func (s *Setting) ReadSection(k string, v interface{}) error {
    err := s.vp.UnmarshalKey(k, v)
    if err != nil {
        return err
    }

    if _, ok := sections[k]; !ok {
        sections[k] = v
    }
    return nil
}

func (s *Setting) ReloadAllSection() error {
    for k, v := range sections {
        err := s.ReadSection(k, v)
        if err != nil {
            return err
        }
    }
    return nil
}
