package proto

import (
    "time"
    "labix.org/v2/mgo/bson"
)

type MetaPack struct {
    Op          string          `json:"op"`
    Data        []DataPack      `json:"data"`
}


type DataPack struct {
    ID          bson.ObjectId   `json:"-" bson:"_id,omitempty"`
    Uid         uint32          `json:"uid" bson:"uid"`
    Name        string          `json:"name" bson:"name"`
    Tsid        uint32          `json:"tsid" bson:"tsid"`
    Num         int64           `json:"num" bson:"num"`
    Time        time.Time       `json:"time" bson:"time"`
}

type    SponsorData    DataPack

type    GiftData       DataPack

