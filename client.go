package dnspod

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	"io"
	"io/ioutil"
	// "log"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

const (
	API_ADDR = "https://dnsapi.cn%s"
)

type Client struct {
	config map[string]string
	l      sync.Mutex
}

func NewClient(id, t string) *Client {
	return &Client{
		config: map[string]string{
			"lang":           "cn",
			"format":         "json",
			"login_token":    id + "," + t,
			"error_on_empty": "yes",
		},
	}
}

func (c *Client) Get(action string, params interface{}, result interface{}) error {
	return c.do("GET", action, params, result)
}
func (c *Client) Post(action string, params interface{}, result interface{}) error {
	return c.do("POST", action, params, result)
}

func (c *Client) do(method string, action string, params interface{}, result interface{}) error {
	apiURL := fmt.Sprintf(API_ADDR, action)
	baseURL, _ := url.Parse(apiURL)

	opts, err := query.Values(params)
	if err != nil {
		return err
	}
	for k, v := range c.config {
		opts.Add(k, v)
	}

	// log.Printf("%s", opts.Encode())
	var body io.Reader

	switch method {
	case "GET":
		body = strings.NewReader(opts.Encode())
	case "POST":
		body = bytes.NewBufferString(opts.Encode())
	}
	hReq, err := http.NewRequest(method, baseURL.String(), body)
	if err != nil {
		return err
	}
	hReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(hReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// log.Printf("%s", respData)
	var commonResponse CommonResponse
	err = json.Unmarshal(respData, &commonResponse)
	if err != nil {
		return err
	}
	if commonResponse.Status.Code != "1" {
		return errors.New(fmt.Sprintf("code:%s ,message:%s", commonResponse.Status.Code, commonResponse.Status.Message))
	}
	err = json.Unmarshal(respData, &result)
	if err != nil {
		return err
	}
	return nil
}
