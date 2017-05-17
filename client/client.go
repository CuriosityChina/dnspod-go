package client

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

	"github.com/CuriosityChina/dnspod-go/logger"
	"github.com/google/go-querystring/query"
	"github.com/qiniu/log"
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

// SetLogLevel set client LogLevel
func (c *Client) SetLogLevel(level string) {
	logger.SetLevel(level)
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
	apiURL := fmt.Sprintf("%s%s", DNSPodAPIBaseURL, action)
	baseURL, err := url.Parse(apiURL)
	if err != nil {
		log.Debug("baseURL parse error: %s", err)
		return err
	}
	logger.Debug("baseURL: %s", baseURL)
	opts, err := query.Values(params)
	if err != nil {
		log.Debug("build query values error: %s", err)
		return err
	}
	logger.Debug("request parammeters: %s", opts.Encode())
	for k, v := range c.config {
		opts.Add(k, v)
	}
	logger.Debug("request parameters with token: %s", opts.Encode())
	var body io.Reader
	switch method {
	case "GET":
		body = strings.NewReader(opts.Encode())
	case "POST":
		body = bytes.NewBufferString(opts.Encode())
	}
	hReq, err := http.NewRequest(method, baseURL.String(), body)
	if err != nil {
		logger.Debug("build http request error: %s", err)
		return err
	}
	hReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(hReq)
	if err != nil {
		logger.Debug("send http request error: %s", err)
		return err
	}
	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Debug("read http response body error: %s", err)
		return err
	}
	var commonResponse CommonResponse
	err = json.Unmarshal(respData, &commonResponse)
	if err != nil {
		logger.Debug("json unmarshal response body error: %s", err)
		return err
	}
	if commonResponse.Status.Code != "1" {
		logger.Debug("code:%s ,message:%s", commonResponse.Status.Code, commonResponse.Status.Message)
		return fmt.Errorf("code:%s ,message:%s", commonResponse.Status.Code, commonResponse.Status.Message)
	}
	return json.Unmarshal(respData, &result)
}
