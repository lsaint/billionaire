package db

import (
    "log"
    "encoding/json"
    "labix.org/v2/mgo"
    "labix.org/v2/mgo/bson"
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

func (this *Dao) GetGiftRank(topn uint32) (ret string) {
    pipe := this.collection.Pipe([]bson.M{
            bson.M{"$group": bson.M{"_id": bson.M{"uid": "$uid", "name": "$name"}, 
                                    "total": bson.M{"$sum": "$num"}}},
            bson.M{"$sort": bson.M{"total": -1}},
            bson.M{"$limit": topn}, 
    })
    iter, items := pipe.Iter(), []bson.M{}
    if err := iter.All(&items); err == nil {
        if js, err := json.Marshal(items); err == nil {
            ret = string(js)
            return
        } else {
            log.Println("Get Marshal err", err)
        }
    } else {
        log.Println("Get All err", err)
    }
    return
}

