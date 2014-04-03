package network

import (
    "fmt"
    "log"
    "time"
    "io/ioutil"
    "net/http"
)

const (
    HTTP_TIME_OUT   =   20
    ADDR            =   ":40404"
)


type HttpReq struct {
    Req             string
    Ret             chan string
}


type HttpSrv struct {
    GetChan             chan *HttpReq
    SaveChan            chan *HttpReq
}

func NewHttpSrv(g, s chan *HttpReq) *HttpSrv {
    return &HttpSrv{g, s}
}

func (this *HttpSrv) Start() {
    http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
            this.get(w, r)
    })
    http.HandleFunc("/save", func(w http.ResponseWriter, r *http.Request) {
            this.save(w, r)
    })

    log.Println("httpsrv runing")
    http.ListenAndServe(ADDR, nil)
}

func (this *HttpSrv) process(c chan *HttpReq, async bool, w http.ResponseWriter, r *http.Request) {
    ret := "done"
    recv_post, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Fprint(w, err.Error())
        return
    }
    req := &HttpReq{string(recv_post), make(chan string)}
    select {
        case c <- req:
            if !async {
                ret = <-req.Ret
            }

        case <-time.After(HTTP_TIME_OUT * time.Second):
            ret = "timeout"
    }
    log.Println("<-->", recv_post, ret)
    fmt.Fprint(w, ret)
}


func (this *HttpSrv) save(w http.ResponseWriter, r *http.Request) {
    this.process(this.SaveChan, true, w, r)
}

func (this *HttpSrv) get(w http.ResponseWriter, r *http.Request) {
    this.process(this.GetChan, false, w, r)
}

