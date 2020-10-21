package conf

import "time"

type ServerSettingS struct {
    RunMode      string
    ServerName         string
    WebServer string
    ServiceServer string
    ReadTimeout  time.Duration
    WriteTimeout time.Duration
}

type EmailSettingS struct {
    Host     string
    Port     int
    UserName string
    Password string
    IsSSL    bool
    From     string
    To       []string
}

type JWTSettingS struct {
    Secret string
    Issuer string
    Expire time.Duration
}

type DatabaseSettingS struct {
    DBType       string
    UserName     string
    Password     string
    Host         string
    DBName       string
    TablePrefix  string
    Charset      string
    ParseTime    bool
    MaxIdleConns int
    MaxOpenConns int
}

type RedisSettings struct {
    Address  []string
}

type TlsSettings struct {
    Addr string
    Cert string
    Key string
}

type LogsSettings struct {
    SaveTime int
    LogDir  string
    LogReduce bool
    LogLevel int
    LogMaxSize int
    LogMaxBackups int
}

type UpLoadSettings struct {
    UploadSavePath        string
    UploadServerUrl       string
    UploadImageMaxSize    int
    UploadImageAllowExts  []string
}