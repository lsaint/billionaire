package db

import (
    "log"
    "labix.org/v2/mgo"
    //"labix.org/v2/mgo/bson"
    "billionaire/conf"
)

const (
)


type Dao struct {
    session         *mgo.Session
    db              *mgo.Database
    collection      *mgo.Collection
}

func NewDao(cl string) *Dao {
    s, err := mgo.Dial(conf.CF.DbAddr)
    if err != nil {
        log.Fatalln(err)
    }
    s.SetMode(mgo.Eventual, true)
    db := s.DB(conf.CF.DbName)
    if err = db.Login(conf.CF.DbUser, conf.CF.DbPw); err != nil {
        log.Fatalln(err)
    }
    collection := db.C(cl)
    return &Dao{s, db, collection}
}

func (this *Dao) Close() {
    this.session.Close()
}

func (this *Dao) Insert(docs ...interface{}) error {
    return this.collection.Insert(docs ...)
}

