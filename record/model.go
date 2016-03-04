package record

type Record struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Line          string `json:"line"`
	Type          string `json:"type"`
	TTL           string `json:"ttl"`
	Value         string `json:"value"`
	Mx            string `json:"mx"`
	Enabled       string `json:"enabled"`
	Status        string `json:"status"`
	MonitorStatus string `json:"monitor_status"`
	Remark        string `json:"remark"`
	UpdatedOn     string `json:"updated_on"`
	UseAqb        string `json:"use_aqb"`
}

type Domain struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Punycode string `json:"punycode"`
	Grade    string `json:"grade"`
	Owner    string `json:"owner"`
}

type Info struct {
	SubDomains  string `json:"sub_domains"`
	RecordTotal string `json:"record_total"`
}
