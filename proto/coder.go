package proto

import (
    "errors"
    "encoding/json"
)

func ParseSaveReq(jn string) (ret *SavePack, err error) {
    ret = &SavePack{}
    if err = json.Unmarshal([]byte(jn), ret); err != nil {
        return
    }
    if ret.Op != "gift" && ret.Op != "sponsor" {
        err = errors.New("op error")
        return
    }
    return
}

func ParseGetReq(jn string) (ret *GetPack, err error) {
    ret = &GetPack{}
    if err = json.Unmarshal([]byte(jn), ret); err != nil {
        return
    }
    if ret.Op != "gift" && ret.Op != "sponsor" {
        err = errors.New("op error")
        return
    }
    if ret.Param < 1 {
        err = errors.New("param error")
        return
    }
    return
}
