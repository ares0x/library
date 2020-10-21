package mongo

import (
    "gopkg.in/mgo.v2"
    "log"
)

func MongoConnect(addr,dataBase string) (*mgo.Session) {
    mongo,err := mgo.Dial(addr)
    if err != nil {
        log.Fatal("Mongo connection failed. Database name: %s,err: %v",addr,err)
    }
    mongo.DB(dataBase)
    return mongo
}