package wx_helper

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"github.com/dchest/uniuri"
	"io/ioutil"
	"log"
	"net/http"
)

//统一下单
func UnifiedOrder(req *UnifiedOrderRequest) (*UnifiedOrderResponse, error) {
	reqData, err := xml.Marshal(req.toXmlRequest())
	log.Println("format ", string(reqData))
	if nil != err {
		log.Println(err)
		return nil, err
	}
	body := bytes.NewBuffer(reqData)
	resp, err := http.Post("https://api.mch.weixin.qq.com/pay/unifiedorder", "application/xml;charset=utf-8", body)
	if nil != err {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		log.Println(err)
		return nil, err
	}
	result := &UnifiedOrderResponse{}
	err = xml.Unmarshal(data, result)
	if nil != err {
		log.Println(err)
		return nil, err
	}
	return result, nil
}

//查询订单
func QueryOrder(outTradeNo string) (*QueryOrderXmlResponse, error) {
	req := &QueryOrderXmlRequest{
		AppId:      FixedParams.AppId,
		MchId:      FixedParams.MchId,
		OutTradeNo: outTradeNo,
		NonceStr:   uniuri.NewLenChars(32, []byte("0123456789abcdefghijklmnopqrstuvwxyz")),
	}
	req.doSign()
	reqData, err := xml.Marshal(req)
	if nil != err {
		log.Println(err)
		return nil, err
	}
	body := bytes.NewBuffer(reqData)
	resp, err := http.Post("https://api.mch.weixin.qq.com/pay/orderquery", "application/xml;charset=utf-8", body)
	if nil != err {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		log.Println(err)
		return nil, err
	}
	result := &QueryOrderXmlResponse{}
	err = xml.Unmarshal(data, result)
	if nil != err {
		log.Println(err)
		return nil, err
	}
	return result, nil
}

//
func DecryptUserInfo(encDataStr, ivStr, sessionKey string) (*WxUserInfo, error) {
	encData, err := base64.StdEncoding.WithPadding(base64.StdPadding).DecodeString(encDataStr)
	if nil != err {
		log.Println(err)
		return nil, err
	}
	aesKey, err := base64.StdEncoding.WithPadding(base64.StdPadding).DecodeString(sessionKey)
	if nil != err {
		log.Println(err)
		return nil, err
	}
	iv, err := base64.StdEncoding.WithPadding(base64.StdPadding).DecodeString(ivStr)
	if nil != err {
		log.Println(err)
		return nil, err
	}
	b, err := aes.NewCipher(aesKey)
	if nil != err {
		log.Println(err)
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(b, iv)
	result := make([]byte, len(encData))
	mode.CryptBlocks(result, encData)
	result = pkcs5UnPadding(result)
	retObj := &WxUserInfo{}
	err = json.Unmarshal(result, retObj)
	return retObj, err
}

/*
func JscodeToSession(jsCode string) (*WxSessionResult, error) {
	params := url.Values{}
	params.Add("appid", _appId)
	params.Add("secret", _appSecret)
	params.Add("js_code", jsCode)
	params.Add("grant_type", "authorization_code")
	resp, err := http.Get("https://api.weixin.qq.com/sns/jscode2session?" + params.Encode())
	if nil != err {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		log.Println(err)
		return nil, err
	}
	result := &WxSessionResult{}
	err = json.Unmarshal(data, result)
	if nil != err {
		log.Println(err)
		return nil, err
	}
	return result, nil
}
*/
/*
func DecryptPhonenumberInfo(encDataStr, ivStr, sessionKey string) (*WxPhoneNumberInfo, error) {
	encData, err := base64.StdEncoding.WithPadding(base64.StdPadding).DecodeString(encDataStr)
	if nil != err {
		log.Println(err)
		return nil, err
	}
	aesKey, err := base64.StdEncoding.WithPadding(base64.StdPadding).DecodeString(sessionKey)
	if nil != err {
		log.Println(err)
		return nil, err
	}
	iv, err := base64.StdEncoding.WithPadding(base64.StdPadding).DecodeString(ivStr)
	if nil != err {
		log.Println(err)
		return nil, err
	}
	b, err := aes.NewCipher(aesKey)
	if nil != err {
		log.Println(err)
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(b, iv)
	result := make([]byte, len(encData))
	mode.CryptBlocks(result, encData)
	result = pkcs5UnPadding(result)
	retObj := &WxPhoneNumberInfo{}
	err = json.Unmarshal(result, retObj)
	return retObj, err
}

// 获取概况趋势
// 累计用户数、转发次数、转发人数
func GetProfileTrend(beginDate string, endDate string) (*ProfileTrend, error) {
	accessToken, err := common.GetWxAccessToken()
	if nil != err {
		log.Println(err)
		return nil, err
	}
	req := &TrendRequest{
		BeginDate: beginDate,
		EndDate:   endDate,
	}
	reqData, err := json.Marshal(req)
	if nil != err {
		log.Println(err)
		return nil, err
	}

	url := "https://api.weixin.qq.com/datacube/getweanalysisappiddailysummarytrend?access_token=" + accessToken
	res, err := http.Post(url, "application/xml;charset=utf-8", bytes.NewBuffer(reqData))
	if nil != err {
		log.Println(err)
		return nil, err
	}
	defer res.Body.Close()
	resData, err := ioutil.ReadAll(res.Body)
	if nil != err {
		log.Println(err)
		return nil, err
	}
	result := &ProfileTrend{}
	err = json.Unmarshal(resData, result)
	if nil != err {
		log.Println(err)
		return nil, err
	}

	return result, err
}

// 获取访问趋势
// 打开次数、访问次数、访问人数、新用户数等
func GetAccessTrend(beginDate string, endDate string) (*AccessTrend, error) {
	accessToken, err := common.GetWxAccessToken()
	if nil != err {
		log.Println(err)
		return nil, err
	}
	req := &TrendRequest{
		BeginDate: beginDate,
		EndDate:   endDate,
	}
	reqData, err := json.Marshal(req)
	if nil != err {
		log.Println(err)
		return nil, err
	}

	url := "https://api.weixin.qq.com/datacube/getweanalysisappiddailyvisittrend?access_token=" + accessToken
	res, err := http.Post(url, "application/xml;charset=utf-8", bytes.NewBuffer(reqData))
	if nil != err {
		log.Println(err)
		return nil, err
	}
	defer res.Body.Close()
	resData, err := ioutil.ReadAll(res.Body)
	if nil != err {
		log.Println(err)
		return nil, err
	}
	result := &AccessTrend{}
	err = json.Unmarshal(resData, result)
	if nil != err {
		log.Println(err)
		return nil, err
	}

	return result, err

}

//获取访问页面的信息
func GetAccessPage(beginDate string, endDate string) (*AccessPage, error) {
	accessToken, err := common.GetWxAccessToken()
	if nil != err {
		log.Println(err)
		return nil, err
	}
	req := &TrendRequest{
		BeginDate: beginDate,
		EndDate:   endDate,
	}
	reqData, err := json.Marshal(req)
	if nil != err {
		log.Println(err)
		return nil, err
	}

	url := "https://api.weixin.qq.com/datacube/getweanalysisappidvisitpage?access_token=" + accessToken
	res, err := http.Post(url, "application/xml;charset=utf-8", bytes.NewBuffer(reqData))
	if nil != err {
		log.Println(err)
		return nil, err
	}
	defer res.Body.Close()
	resData, err := ioutil.ReadAll(res.Body)
	if nil != err {
		log.Println(err)
		return nil, err
	}
	result := &AccessPage{}
	err = json.Unmarshal(resData, result)
	if nil != err {
		log.Println(err)
		return nil, err
	}

	return result, err
}
*/
