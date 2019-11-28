package wx_helper

type TrendRequest struct {
	BeginDate string `json:"begin_date" desc:"开始日期"`
	EndDate   string `json:"end_date" desc:"结束日期"`
}

// 概况趋势
type ProfileTrend struct {
	List []ProfileTrendItem `json:"list" desc:"list"`
}

type ProfileTrendItem struct {
	RefDate    string `json:"ref_date" desc:"时间"`
	VisitTotal int    `json:"visit_total" desc:"累计用户数"`
	SharePv    int    `json:"share_pv" desc:"转发次数"`
	ShareUv    int    `json:"share_uv" desc:"转发人数"`
}

//访问趋势
type AccessTrend struct {
	List []AccessTrendItem `json:"list" desc:"list"`
}

type AccessTrendItem struct {
	RefDate         string  `json:"ref_date" desc:"时间"`
	SessionCnt      int     `json:"session_cnt" desc:"打开次数"`
	VisitPv         int     `json:"visit_pv" desc:"访问次数"`
	VisitUv         int     `json:"visit_uv" desc:"访问人数"`
	VisitUvNew      int     `json:"visit_uv_new" desc:"新用户数"`
	StayTimeUv      float64 `json:"stay_time_uv" desc:"人均停留时长，单位秒"`
	StayTimeSession float64 `json:"stay_time_session" desc:"次均停留时长，单位秒"`
}

//访问页面
type AccessPage struct {
	RefDate string           `json:"ref_date" desc:"时间"`
	List    []AccessPageItem `json:"list" desc:"列表"`
}

type AccessPageItem struct {
	PagePath       string  `json:"page_path" desc:"页面路径"`
	PageVisitPv    int     `json:"page_visit_pv" desc:"访问次数"`
	PageVisitUv    int     `json:"page_visit_uv" desc:"访问人数"`
	PageStaytimePv float64 `json:"page_staytime_pv" desc:"次均停留时长"`
	EntrypagePv    int     `json:"entrypage_pv" desc:"进入页次数"`
	ExitpagePv     int     `json:"exitpage_pv" desc:"退出页次数"`
	PageSharePv    int     `json:"page_share_pv" desc:"转发次数"`
	PageShareUv    int     `json:"page_share_uv" desc:"转发人数"`
}
