package mysql

type DbSetting struct {
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

