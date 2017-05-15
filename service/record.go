package service

import (
	"github.com/CuriosityChina/dnspod-go/client"
)

// RecordService dnspod record service
type RecordService struct {
	*client.Client
}

// NewRecordService return a new RecordService
func NewRecordService(clt *client.Client) *RecordService {
	return &RecordService{
		Client: clt,
	}
}

// RecordCreateRequest record create input
type RecordCreateRequest struct {
	DomainID   int    `url:"domain_id"`
	SubDomain  string `url:"sub_domain"`
	RecordType string `url:"record_type"`
	RecordLine string `url:"record_line"`
	Value      string `url:"value"`
	MX         string `url:"mx"`
	TTL        string `url:"ttl"`
	Status     string `url:"status"`
	Weight     string `url:"weight"`
}

// RecordCreateResponse record create output
type RecordCreateResponse struct {
	Record struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Status string `json:"status"`
	} `json:"record"`
	*client.CommonResponse
}

// RecordCreate record create action
func (c *RecordService) RecordCreate(req RecordCreateRequest) (RecordCreateResponse, error) {
	var result RecordCreateResponse
	err := c.Post("/Record.Create", &req, &result)
	return result, err
}

// RecordListRequest record list input
type RecordListRequest struct {
	DomainID  int    `url:"domain_id"`
	Offset    int    `url:"offset"`
	Length    int    `url:"length"`
	SubDomain string `url:"sub_domain"`
	Keyword   string `url:"keyword"`
}

// RecordListResponse record list output
type RecordListResponse struct {
	Domain  ListDomainType   `json:"domain"`
	Info    ListInfoType     `json:"info"`
	Records []ListRecordType `json:"records"`
	*client.CommonResponse
}

// RecordList record list action
func (c *RecordService) RecordList(req RecordListRequest) (RecordListResponse, error) {
	var result RecordListResponse
	err := c.Post("/Record.List", &req, &result)
	return result, err
}

// RecordModifyRequest record modify input
type RecordModifyRequest struct {
	DomainID   int    `url:"domain_id"`
	RecordID   string `url:"record_id"`
	SubDomain  string `url:"sub_domain"`
	RecordType string `url:"record_type"`
	RecordLine string `url:"record_line"`
	Value      string `url:"value"`
	MX         string `url:"mx"`
	TTL        string `url:"ttl"`
	Status     string `url:"status"`
	Weight     string `url:"weight"`
}

// RecordModifyResponse record modify output
type RecordModifyResponse struct {
	Record struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Value string `json:"value"`
		Staus string `json:"status"`
	} `json:"record"`
	*client.CommonResponse
}

// RecordModify record modify action
func (c *RecordService) RecordModify(req RecordModifyRequest) (RecordModifyResponse, error) {
	var result RecordModifyResponse
	err := c.Post("/Record.Modify", &req, &result)
	return result, err
}

// RecordInfoRequest record info input
type RecordInfoRequest struct {
	DomainID int    `url:"domain_id"`
	RecordID string `url:"record_id"`
}

// RecordInfoResponse record info output
type RecordInfoResponse struct {
	Domain InfoDomainType `json:"domain"`
	Record InfoRecordType `json:"record"`
	*client.CommonResponse
}

// RecordInfo record info output
func (c *RecordService) RecordInfo(req RecordInfoRequest) (RecordInfoResponse, error) {
	var result RecordInfoResponse
	err := c.Post("/Record.Info", &req, &result)
	return result, err
}

// RecordRemoveRequest record remove input
type RecordRemoveRequest struct {
	DomainID int    `url:"domain_id"`
	RecordID string `url:"record_id"`
}

// RecordRemoveResponse record remove output
type RecordRemoveResponse *client.CommonResponse

// RecordRemove record remove action
func (c *RecordService) RecordRemove(req RecordRemoveRequest) (RecordRemoveResponse, error) {
	var result RecordRemoveResponse
	err := c.Post("/Record.Remove", &req, &result)
	return result, err
}
