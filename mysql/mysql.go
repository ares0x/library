package mysql

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "log"
    "time"
)

var Db *gorm.DB

func openDB(username, password, addr, name string,maxIdlConns, maxOpenConns int) *gorm.DB {
    config := fmt.Sprintf(username+":"+password+"@tcp("+addr+")/"+name+"?charset=utf8&parseTime=true")
    log.Println("config:",config)
    db, err := gorm.Open("mysql", config)
    if err != nil {
        log.Fatal("Database connection failed. Database name: %s,err: %v",name,err)
    }
    db.DB().SetMaxIdleConns(maxIdlConns)
    db.DB().SetMaxOpenConns(maxOpenConns)
    db.DB().SetConnMaxLifetime(time.Second * 300)
    return db
}

func InitSql(username, password, addr, name string,maxIdlConns, maxOpenConns int) {
    Db = openDB(username, password, addr, name ,maxIdlConns, maxOpenConns)
}

func GetDB() *gorm.DB {
    return Db
}