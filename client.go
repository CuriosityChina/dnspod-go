package dnspod

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/google/go-querystring/query"
)

const (
	// DNSPodAPIBaseURL dnspod api base url
	DNSPodAPIBaseURL = "https://dnsapi.cn"
)

// Client dnspod client
type Client struct {
	config map[string]string
	lock   sync.Mutex
}

// NewClient return a new dnspod client
func NewClient(id, token string) *Client {
	return &Client{
		config: map[string]string{
			"lang":           "cn",
			"format":         "json",
			"login_token":    id + "," + token,
			"error_on_empty": "yes",
		},
	}
}

// Get HTTP GET
func (c *Client) Get(action string, params interface{}, result interface{}) error {
	return c.do("GET", action, params, result)
}

// Post HTTP POST
func (c *Client) Post(action string, params interface{}, result interface{}) error {
	return c.do("POST", action, params, result)
}

// Post HTTP Do
func (c *Client) do(method string, action string, params interface{}, result interface{}) error {
	apiURL := fmt.Sprintf("%s/%s", DNSPodAPIBaseURL, action)
	baseURL, _ := url.Parse(apiURL)
	opts, err := query.Values(params)
	if err != nil {
		return err
	}
	for k, v := range c.config {
		opts.Add(k, v)
	}
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
	var commonResponse CommonResponse
	err = json.Unmarshal(respData, &commonResponse)
	if err != nil {
		return err
	}
	if commonResponse.Status.Code != "1" {
		return fmt.Errorf("code:%s ,message:%s", commonResponse.Status.Code, commonResponse.Status.Message)
	}
	return json.Unmarshal(respData, &result)
}
