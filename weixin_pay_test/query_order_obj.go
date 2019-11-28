package wx_helper

import (
	"encoding/xml"
	"fmt"
	"reflect"
)

type QueryOrderXmlRequest struct {
	XMLName    xml.Name `xml:"xml"`          //不用填写
	AppId      string   `xml:"appid"`        //不用填写
	MchId      string   `xml:"mch_id"`       //不用填写
	OutTradeNo string   `xml:"out_trade_no"` //商户订单号
	NonceStr   string   `xml:"nonce_str"`    //随机字符串
	Sign       string   `xml:"sign"`         //签名
}

func (this *QueryOrderXmlRequest) doSign() {
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

type QueryOrderXmlResponse struct {
	ReturnCode     string `xml:"return_code"`      //返回状态码
	ReturnMsg      string `xml:"return_msg"`       //返回信息
	AppId          string `xml:"appid"`            //小程序ID
	MchId          string `xml:"mch_id"`           //商户号
	NonceStr       string `xml:"nonce_str"`        //随机字符串
	Sign           string `xml:"sign"`             //签名
	ResultCode     string `xml:"result_code"`      //业务结果
	ErrCode        string `xml:"err_code"`         //错误代码
	ErrCodeDes     string `xml:"err_code_des"`     //错误代码描述
	OpenId         string `xml:"openid"`           //openId
	TradeType      string `xml:"trade_type"`       //交易类型
	TradeState     string `xml:"trade_state"`      //交易状态
	BankType       string `xml:"bank_type"`        //付款银行
	TotalFee       uint64 `xml:"total_fee"`        //订单金额
	CashFee        uint64 `xml:"cash_fee"`         //现金支付金额
	TransactionId  string `xml:"transaction_id"`   //微信支付订单号
	OutTradeNo     string `xml:"out_trade_no"`     //商户订单号
	Attach         string `xml:"attach"`           //附加数据
	TimeEnd        string `xml:"time_end"`         //支付完成时间
	TradeStateDesc string `xml:"trade_state_desc"` //交易状态描述
}
