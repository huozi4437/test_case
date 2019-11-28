package wx_helper

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"time"

	"github.com/dchest/uniuri"
)

type WxSessionResult struct {
	Errcode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
}

type WxPaymentRequest struct {
	AppId     string `json:"appId"`
	TimeStamp int64  `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}

func NewWxPaymentRequest(prePayId string) *WxPaymentRequest {
	ret := &WxPaymentRequest{
		AppId:     FixedParams.AppId,
		TimeStamp: time.Now().Unix(),
		NonceStr:  uniuri.NewLenChars(32, []byte("0123456789abcdefghijklmnopqrstuvwxyz")),
		Package:   "prepay_id=" + prePayId,
		SignType:  "MD5",
	}
	objValue := reflect.ValueOf(*ret)
	objSt := reflect.TypeOf(*ret)

	attrMap := map[string]string{}
	for i := 0; i < objValue.NumField(); i++ {
		field := objValue.Field(i)
		t := objSt.Field(i)
		tagName := t.Tag.Get("json")
		attrMap[tagName] = fmt.Sprint(field)
	}
	ret.PaySign = Sign(attrMap)
	return ret
}

func Sign(attrMap map[string]string) string {
	attrKeyList := []string{}
	for key, _ := range attrMap {
		attrKeyList = append(attrKeyList, key)
	}
	sort.Strings(attrKeyList)
	pairList := []string{}
	for _, attrKey := range attrKeyList {
		value := attrMap[attrKey]
		if "" == value {
			continue
		}
		pairList = append(pairList, fmt.Sprintf("%s=%s", attrKey, value))
	}
	pairList = append(pairList, "key="+FixedParams.MchKey)
	toSign := strings.Join(pairList, "&")
	h := md5.New()
	h.Write([]byte(toSign))
	sign := hex.EncodeToString(h.Sum(nil))
	return strings.ToUpper(sign)
}
