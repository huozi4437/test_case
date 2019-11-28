package wx_helper

type WxUserInfoWaterMark struct {
	AppId string `json:"appid"`
}

type WxUserInfo struct {
	OpenId    string              `json:"openId"`
	NickName  string              `json:"nickName"`
	Gender    int                 `json:"gender"`
	City      string              `json:"city"`
	Province  string              `json:"province"`
	Country   string              `json:"country"`
	AvatarUrl string              `json:"avatarUrl"`
	WaterMark WxUserInfoWaterMark `json:"watermark"`
}
