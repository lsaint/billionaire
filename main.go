// L'billionaire

package main


import (
    "os"
    "log"
    "syscall"
    "os/signal"

    "billionaire/network"
)

func handleSig() {
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)

    for sig := range c {
        log.Printf("__handle__signal__ %v", sig)
        return
    }
}

func main() {
    getchan, savechan := make(chan *network.HttpReq, 1024), make(chan *network.HttpReq, 10240)
    httpsrv := network.NewHttpSrv(getchan, savechan)
    go httpsrv.Start()
    mgr := NewBillboardMgr(getchan, savechan)
    go mgr.Start()

    handleSig()
}

