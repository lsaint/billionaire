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
