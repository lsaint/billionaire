package main


import (
    "log"
    "billionaire/network"
    "billionaire/proto"
    "billionaire/db"
)


type BillboardMgr struct {
    getChan             chan *network.HttpReq
    saveChan            chan *network.HttpReq

    giftDao             *db.Dao
    sponsorDao          *db.Dao
    name2dao            map[string]*db.Dao
}

func NewBillboardMgr(g, s chan *network.HttpReq) *BillboardMgr {
    giftDao := db.NewDao("gift")
    sponsorDao := db.NewDao("sponsor")
    dt := map[string]*db.Dao{"gift": giftDao, 
                            "sponsor": sponsorDao}
    return &BillboardMgr{getChan: g, 
                        saveChan: s, 
                        giftDao: giftDao,
                        sponsorDao: sponsorDao,
                        name2dao: dt}
}

func (this *BillboardMgr) Start() {
    log.Println("BillboardMgr running")
    go func() {
        for { select {
            case req := <-this.getChan:
                this.onGet(req)
        }}
    }()

    go func() {
        for { select {
            case req := <-this.saveChan:
                this.onSave(req)
        }}
    }()
}

func (this *BillboardMgr) onGet(req *network.HttpReq) {
    if pack, err := proto.ParseGetReq(req.Req); err == nil {
        if dao, exist := this.name2dao[pack.Op]; exist {
            req.Ret <- dao.GetGiftRank(pack.Param)
        } else {
            log.Println("not exist op", pack.Op)        
        }
    } else {
        req.Ret <- err.Error()
    }
}


func (this *BillboardMgr) onSave(req *network.HttpReq) {
    if pack, err := proto.ParseSaveReq(req.Req); err == nil {
        if dao, exist := this.name2dao[pack.Op]; exist {
            for _, item := range pack.Data {
                dao.Insert(&item)
            }
        } else {
            log.Println("not exist op", pack.Op)        
        }
    } else {
        log.Println("onSave err", err)        
    }
}


