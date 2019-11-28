package wx_helper

import (
	"encoding/xml"
	"fmt"
	"reflect"
)

var FixedParams *FixedParamsStc

//统一下单api数据
type UnifiedOrderXmlRequest struct {
	XMLName        xml.Name `xml:"xml"`              //不用填写
	AppId          string   `xml:"appid"`            //公众账号ID，不用填写
	MchId          string   `xml:"mch_id"`           //商户号，不用填写
	NonceStr       string   `xml:"nonce_str"`        //随机字符串，不长于32位
	Sign           string   `xml:"sign"`             //签名，md5类型
	Body           string   `xml:"body"`             //商品描述
	Attach         string   `xml:"attach"`           //附加数据
	OutTradeNo     string   `xml:"out_trade_no"`     //商户订单号,32个字符内
	TotalFee       int64    `xml:"total_fee"`        //订单总金额，单位为分
	SpbillCreateIp string   `xml:"spbill_create_ip"` //终端IP，用户的客户端IP
	NotifyUrl      string   `xml:"notify_url"`       //接收微信支付异步通知回调地址，通知url必须为直接可访问的url，不能携带参数。
	TradeType      string   `xml:"trade_type"`       //交易类型, 取值如下：NATIVE
	ProductId      string   `xml:"product_id"`       //商品ID，NATIVE模式必传
}

//固定配置项
type FixedParamsStc struct {
	AppId     string //公众账号ID
	AppSecret string //秘钥
	MchId     string //商户号
	MchKey    string //秘钥
	NotifyUrl string //接收微信支付异步通知回调地址，通知url必须为直接可访问的url，不能携带参数。
	TradeType string //交易类型, 取值如下：NATIVE
}

func InitParams(params *FixedParamsStc) {
	FixedParams = params
	FixedParams.TradeType = "NATIVE"
}

//调用时传入项
type UnifiedOrderRequest struct {
	NonceStr       string //随机字符串，不长于32位
	Body           string //商品描述
	Attach         string //附加数据
	OutTradeNo     string //商户系统内部的订单号, 32个字符内
	TotalFee       int64  //订单总金额，单位为分
	SpbillCreateIp string //终端IP
	ProductId      string //商品ID
}

func (this *UnifiedOrderRequest) toXmlRequest() *UnifiedOrderXmlRequest {
	req := &UnifiedOrderXmlRequest{}
	req.NonceStr = this.NonceStr
	req.Body = this.Body
	req.Attach = this.Attach
	req.OutTradeNo = this.OutTradeNo
	req.TotalFee = this.TotalFee
	req.SpbillCreateIp = this.SpbillCreateIp
	req.TradeType = FixedParams.TradeType
	req.AppId = FixedParams.AppId
	req.MchId = FixedParams.MchId
	req.NotifyUrl = FixedParams.NotifyUrl
	req.doSign()
	return req
}

func (this *UnifiedOrderXmlRequest) doSign() {
	objValue := reflect.ValueOf(*this)
	objSt := reflect.TypeOf(*this)
	attrMap := map[string]string{}
	for i := 0; i < objValue.NumField(); i++ {
		field := objValue.Field(i)
		t := objSt.Field(i)
		tagName := t.Tag.Get("xml")
		if "xml" == tagName {
			continue
		}
		attrMap[tagName] = fmt.Sprint(field)
	}
	this.Sign = Sign(attrMap)
}

type UnifiedOrderResponse struct {
	ReturnCode string `xml:"return_code"`  //返回状态码
	ReturnMsg  string `xml:"return_msg"`   //返回信息
	AppId      string `xml:"appid"`        //公众账号ID
	MchId      string `xml:"mch_id"`       //商户号
	NonceStr   string `xml:"nonce_str"`    //随机字符串，不长于32位
	Sign       string `xml:"sign"`         //签名
	ResultCode string `xml:"result_code"`  //业务结果
	ErrCode    string `xml:"err_code"`     //错误码
	ErrCodeDes string `xml:"err_code_des"` //错误描述
	TradeType  string `xml:"trade_type"`   //取值如下：NATIVE
	PrepayId   string `xml:"prepay_id"`    //预支付交易会话标识
	CodeUrl    string `xml:"code_url"`     //二维码链接
}
