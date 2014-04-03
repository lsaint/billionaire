package proto


type MetaPack struct {
    Op          string          `json:"op"`
    Data        []DataPack      `json:"data"`
}


type DataPack struct {
    Uid         uint32          `json:"uid"`
    Name        string          `json:"name"`
    Tsid        uint32          `json:"tsid"`
    Num         int64           `json:"num"`
    Time        uint32          `json:"time"`
}

type    SponsorData    DataPack

type    GiftData       DataPack

