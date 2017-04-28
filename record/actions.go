package record

import (
	"github.com/CuriosityChina/dnspod-go"
)

type RECORD struct {
	*dnspod.Client
}

func NewClient(clt *dnspod.Client) *RECORD {
	return &RECORD{
		Client: clt,
	}
}

type RecordCreateRequest struct {
	DomainId   int    `url:"domain_id"`
	SubDomain  string `url:"sub_domain"`
	RecordType string `url:"record_type"`
	RecordLine string `url:"record_line"`
	Value      string `url:"value"`
	MX         string `url:"mx"`
	TTL        string `url:"ttl"`
	Status     string `url:"status"`
	Weight     string `url:"weight"`
}
type RecordCreateResponse struct {
	Record struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Value string `json:"value"`
		Staus string `json:"status"`
	} `json:"record"`
	*dnspod.CommonResponse
}

func (c *RECORD) RecordCreate(req RecordCreateRequest) (RecordCreateResponse, error) {
	var result RecordCreateResponse
	err := c.Post("/Record.Create", &req, &result)
	return result, err
}

type RecordListRequest struct {
	DomainId  int    `url:"domain_id"`
	Offset    int    `url:"offset"`
	Length    int    `url:"length"`
	SubDomain string `url:"sub_domain"`
	Keyword   string `url:"keyword"`
}
type RecordListResponse struct {
	Domain  Domain   `json:"domain"`
	Info    Info     `json:"info"`
	Records []Record `json:"records"`
	*dnspod.CommonResponse
}

func (c *RECORD) RecordList(req RecordListRequest) (RecordListResponse, error) {
	var result RecordListResponse
	err := c.Post("/Record.List", &req, &result)
	return result, err
}

type RecordModifyRequest struct {
	DomainId   int    `url:"domain_id"`
	RecordId   string `url:"record_id"`
	SubDomain  string `url:"sub_domain"`
	RecordType string `url:"record_type"`
	RecordLine string `url:"record_line"`
	Value      string `url:"value"`
	MX         string `url:"mx"`
	TTL        string `url:"ttl"`
	Status     string `url:"status"`
	Weight     string `url:"weight"`
}
type RecordModifyResponse struct {
	Domain Domain `json:"domain"`
	Record struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Value string `json:"value"`
		Staus string `json:"status"`
	} `json:"record"`
	*dnspod.CommonResponse
}

func (c *RECORD) RecordModify(req RecordModifyRequest) (RecordModifyResponse, error) {
	var result RecordModifyResponse
	err := c.Post("/Record.Modify", &req, &result)
	return result, err
}

type RecordInfoRequest struct {
	DomainId int    `url:"domain_id"`
	RecordId string `url:"record_id"`
}
type RecordInfoResponse struct {
	Domain Domain `json:"domain"`
	Record Record `json:"record"`
	*dnspod.CommonResponse
}

func (c *RECORD) RecordInfo(req RecordInfoRequest) (RecordInfoResponse, error) {
	var result RecordInfoResponse
	err := c.Post("/Record.Info", &req, &result)
	return result, err
}

type RecordRemoveRequest struct {
	DomainId int    `url:"domain_id"`
	RecordId string `url:"record_id"`
}
type RecordRemoveResponse *dnspod.CommonResponse

func (c *RECORD) RecordRemove(req RecordRemoveRequest) (RecordRemoveResponse, error) {
	var result RecordRemoveResponse
	err := c.Post("/Record.Remove", &req, &result)
	return result, err
}
