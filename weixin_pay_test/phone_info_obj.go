package wx_helper

type WxPhoneWatermark struct {
	Appid     string `json:"appid"`
	Timestamp int64  `json:"timestamp"`
}

type WxPhoneNumberInfo struct {
	PhoneNumber     string           `json:"phoneNumber"`
	PurePhoneNumber string           `json:"purePhoneNumber"`
	CountryCode     string           `json:"countryCode"`
	Watermark       WxPhoneWatermark `json:"watermark"`
}
