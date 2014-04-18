package db

import (
    "log"
    "time"
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

    today           time.Time
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
    dao := &Dao{session: s, db:db, collection:collection}
    dao.gentoday()
    return dao
}

func (this *Dao) Close() {
    this.session.Close()
}

func (this *Dao) Insert(docs ...interface{}) error {
    return this.collection.Insert(docs ...)
}

func (this *Dao) GetGiftRank(topn uint32) (ret string) {
    pipe := this.collection.Pipe([]bson.M{
            bson.M{"$match": bson.M{"time": bson.M{"$gt": this.Today()}}},
            bson.M{"$group": bson.M{"_id": "$uid", 
                                    "name": bson.M{"$last": "$name"},
                                    "total": bson.M{"$sum": "$num"},
                                    "time": bson.M{"$last": "$time"}}},
            bson.M{"$sort": bson.M{"total": -1, "time": -1}},
            bson.M{"$limit": topn}, 
            bson.M{"$project": bson.M{"_id":1, "name":1, "total":1}}, 
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

func (this *Dao) Today() time.Time {
    if time.Now().Day() > this.today.Day() {
        this.gentoday()
    }
    return this.today
}

func (this *Dao) gentoday() {
    now := time.Now()
    this.today = time.Date(now.Year(), now.Month(), now.Day(), 0,0,0,0, now.Location())
}

