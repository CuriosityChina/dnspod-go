package dnspod

import (
	"encoding/json"
	"github.com/magicshui/dnspod-go/domain"
)

func (c *Client) DomainList(dtype string, offset int, length int, group_id string, keyword string) (domain.DomainList, error) {
	var request = map[string]interface{}{"type": dtype, "offset": offset, "length": length, "group_id": group_id, "keyword": keyword}
	var result domain.DomainList
	raw_result, err := c.Post("Domain.List", request)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(raw_result, &result)
	if err != nil {
		return result, err
	}
	return result, nil

}
