package mongo

import (
    "bytes"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "io"
    "log"
)

var mongoDb *mgo.Session

func InitMongo(addr,dataBase string)  {
    mongo,err := mgo.Dial(addr)
    if err != nil {
        log.Fatal("Mongo connection failed. Database name: %s,err: %v",addr,err)
    }
    mongo.DB(dataBase)
    mongoDb = mongo
}

func QueryFileById(_id string) ([]byte, error) {
    session, err := otc_mongo.MongoSession()
    if err != nil {
        otc_mongo.MongoDBLogger.Error(fmt.Sprintf("create mongo session error: %v", err))
        return nil, err
    }
    gridFsReader, err := session.DB("file").GridFS("fs").OpenId(bson.ObjectIdHex(_id))
    if err != nil {
        otc_mongo.MongoDBLogger.Error(fmt.Sprintf("create mongo gridFsReader error: %v", err))
        return nil, err
    }
    defer gridFsReader.Close()

    buffer := bytes.NewBuffer(nil)
    _, err = io.Copy(buffer, gridFsReader)
    if err != nil {
        otc_mongo.MongoDBLogger.Error(fmt.Sprintf("mongo clinet io Copy error: %v", err))
        return nil, err
    }

    return buffer.Bytes(), nil
}