package conf

import (
    "log"
    "io/ioutil"
    "encoding/json"
)

var CF *JsonConfig
func init() {
    CF = &JsonConfig{}
    CF.ReadConfig()
}


type JsonConfig struct {
    DbAddr          string      
    DbName          string      
    DbUser          string     
    DbPw            string      
}

func (this *JsonConfig) ReadConfig() {
    b, err := ioutil.ReadFile("./conf/config.txt") 
    if err != nil {
        log.Fatalln(err)
    }
    err = json.Unmarshal(b, this)
    if err != nil {
        log.Fatalln(err)
    }

    log.Printf("%+v\n", *this)
    log.Printf("Read Config Done")
}


