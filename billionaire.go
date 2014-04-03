
package main


import (
    "log"
    "billionaire/network"
)


type BillboardMgr struct {
    getChan         chan *network.HttpReq
    saveChan        chan *network.HttpReq
}

func NewBillboardMgr(g, s chan *network.HttpReq) *BillboardMgr {
    bb := &BillboardMgr{g, s}
    return bb
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
}


func (this *BillboardMgr) onSave(req *network.HttpReq) {
}

