package dnspod

import (
	"encoding/json"
	"github.com/magicshui/dnspod-go/record"
)

func (c *Client) RecordList(domain_id int, offset int, length int, sub_domain string, keyword string) (record.RecordList, error) {
	var request = map[string]interface{}{"domain_id": domain_id, "offset": offset, "length": length, "sub_domain": sub_domain, "keyword": keyword}
	var result record.RecordList
	raw_result, err := c.Post("Record.List", request)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(raw_result, &result)
	if err != nil {
		return result, err
	}
	return result, nil

}
func (c *Client) RecordCreate(domain_id int, sub_domain string, record_type string, record_line string, value string, mx int, ttl int64) (record.RecordCreate, error) {
	var request = map[string]interface{}{"domain_id": domain_id, "sub_domain": sub_domain, "value": value, "record_type": record_type, "record_line": record_line, "mx": mx, "ttl": ttl}
	var result record.RecordCreate
	raw_result, err := c.Post("Record.Create", request)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(raw_result, &result)
	if err != nil {
		return result, err
	}
	return result, nil

}
