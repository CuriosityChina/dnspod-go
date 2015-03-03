package dnspod

import (
	"io/ioutil"
	"net/http"
	neturl "net/url"
	"reflect"
	"strconv"
)

type Client struct {
	LoginEmail    string `json:"login_email"`
	LoginPassword string `json:"login_password"`
	Format        string `json:"format"`
	Lang          string `json:"lang"`
	ErrorOnEmpty  string `json:"error_on_empty"`
	UserId        string `json:"user_id"`
}

func (c *Client) Post(endpoint string, request interface{}) ([]byte, error) {

	var _request = request.(map[string]interface{})
	_request["login_email"] = c.LoginEmail
	_request["login_password"] = c.LoginPassword
	_request["format"] = c.Format
	_request["lang"] = c.Lang
	_request["error_on_empty"] = c.ErrorOnEmpty
	_request["user_id"] = c.UserId

	url := "https://dnsapi.cn/" + endpoint

	data := neturl.Values{}

	for k, v := range _request {
		sv := reflect.TypeOf(_request[k]).String()
		switch sv {
		case "string":
			data.Add(k, v.(string))
		case "int":
			data.Add(k, strconv.Itoa(v.(int)))
		case "int64":

			data.Add(k, strconv.FormatInt(v.(int64), 10))
		}

	}
	res, err := http.PostForm(url, data)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte(""), err
	}
	return result, err

}
func (c *Client) Get() {

}
