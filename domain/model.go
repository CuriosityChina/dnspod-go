package domain

type Domain struct {
	AuthToAnquanbao  bool    `json:"auth_to_anquanbao"`
	CnameSpeedup     string  `json:"cname_speedup"`
	CreatedOn        string  `json:"created_on"`
	ExtStatus        string  `json:"ext_status"`
	Grade            string  `json:"grade"`
	GradeTitle       string  `json:"grade_title"`
	GroupID          string  `json:"group_id"`
	ID               float64 `json:"id"`
	IsMark           string  `json:"is_mark"`
	IsVip            string  `json:"is_vip"`
	Name             string  `json:"name"`
	Owner            string  `json:"owner"`
	Punycode         string  `json:"punycode"`
	Records          string  `json:"records"`
	Remark           string  `json:"remark"`
	SearchenginePush string  `json:"searchengine_push"`
	Status           string  `json:"status"`
	Ttl              string  `json:"ttl"`
	UpdatedOn        string  `json:"updated_on"`
}

type DomainList struct {
	Domains []struct {
		AuthToAnquanbao  bool    `json:"auth_to_anquanbao"`
		CnameSpeedup     string  `json:"cname_speedup"`
		CreatedOn        string  `json:"created_on"`
		ExtStatus        string  `json:"ext_status"`
		Grade            string  `json:"grade"`
		GradeTitle       string  `json:"grade_title"`
		GroupID          string  `json:"group_id"`
		ID               float64 `json:"id"`
		IsMark           string  `json:"is_mark"`
		IsVip            string  `json:"is_vip"`
		Name             string  `json:"name"`
		Owner            string  `json:"owner"`
		Punycode         string  `json:"punycode"`
		Records          string  `json:"records"`
		Remark           string  `json:"remark"`
		SearchenginePush string  `json:"searchengine_push"`
		Status           string  `json:"status"`
		Ttl              string  `json:"ttl"`
		UpdatedOn        string  `json:"updated_on"`
	} `json:"domains"`
	Info struct {
		AllTotal      int `json:"all_total"`
		DomainTotal   int `json:"domain_total"`
		ErrorTotal    int `json:"error_total"`
		IsmarkTotal   int `json:"ismark_total"`
		LockTotal     int `json:"lock_total"`
		MineTotal     int `json:"mine_total"`
		PauseTotal    int `json:"pause_total"`
		ShareOutTotal int `json:"share_out_total"`
		ShareTotal    int `json:"share_total"`
		SpamTotal     int `json:"spam_total"`
		VipExpire     int `json:"vip_expire"`
		VipTotal      int `json:"vip_total"`
	} `json:"info"`
	Status struct {
		Code      string `json:"code"`
		CreatedAt string `json:"created_at"`
		Message   string `json:"message"`
	} `json:"status"`
}
