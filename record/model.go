package record

type RecordList struct {
	Domain struct {
		Grade    string  `json:"grade"`
		ID       float64 `json:"id"`
		Name     string  `json:"name"`
		Owner    string  `json:"owner"`
		Punycode string  `json:"punycode"`
	} `json:"domain"`
	Info struct {
		RecordTotal string `json:"record_total"`
		SubDomains  string `json:"sub_domains"`
	} `json:"info"`
	Records []struct {
		Enabled       string `json:"enabled"`
		ID            string `json:"id"`
		Line          string `json:"line"`
		MonitorStatus string `json:"monitor_status"`
		Mx            string `json:"mx"`
		Name          string `json:"name"`
		Remark        string `json:"remark"`
		Status        string `json:"status"`
		Ttl           string `json:"ttl"`
		Type          string `json:"type"`
		UpdatedOn     string `json:"updated_on"`
		UseAqb        string `json:"use_aqb"`
		Value         string `json:"value"`
	} `json:"records"`
	Status struct {
		Code      string `json:"code"`
		CreatedAt string `json:"created_at"`
		Message   string `json:"message"`
	} `json:"status"`
}

type RecordCreate struct {
	Record struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Status string `json:"status"`
	} `json:"record"`
	Status struct {
		Code      string `json:"code"`
		CreatedAt string `json:"created_at"`
		Message   string `json:"message"`
	} `json:"status"`
}
