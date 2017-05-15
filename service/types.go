package service

// ListRecordType record list record type
type ListRecordType struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Line          string `json:"line"`
	Type          string `json:"type"`
	TTL           string `json:"ttl"`
	Value         string `json:"value"`
	MX            string `json:"mx"`
	Enabled       string `json:"enabled"`
	Status        string `json:"status"`
	MonitorStatus string `json:"monitor_status"`
	Remark        string `json:"remark"`
	UpdatedOn     string `json:"updated_on"`
	UseAqb        string `json:"use_aqb"`
	SubDomain     string `json:"sub_domain"`
	Weight        string `json:"weight"`
}

// ListDomainType record list domain type
type ListDomainType struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Punycode  string `json:"punycode"`
	Grade     string `json:"grade"`
	Owner     string `json:"owner"`
	ExtStatus string `json:"dnserror"`
	TTL       int    `json:"ttl"`
}

// ListInfoType record list info type
type ListInfoType struct {
	SubDomains  string `json:"sub_domains"`
	RecordTotal string `json:"record_total"`
}

// InfoDomainType record info domain type
type InfoDomainType struct {
	ID          int    `json:"id"`
	Domain      string `json:"domain"`
	DomainGrade string `json:"domain_grade"`
}

// InfoRecordType record info record type
type InfoRecordType struct {
	ID            string `json:"id"`
	SubDomain     string `json:"sub_domain"`
	RecordType    string `json:"record_type"`
	RecordLine    string `json:"record_line"`
	RecordLineID  string `json:"record_line_id"`
	Value         string `json:"value"`
	Weight        string `json:"weight"`
	MX            string `json:"mx"`
	TTL           string `json:"ttl"`
	Enabled       string `json:"enabled"`
	MonitorStatus string `json:"monitor_status"`
	Remark        string `json:"remark"`
	UpdatedOn     string `json:"updated_on"`
	DomainID      string `json:"domain_id"`
}
