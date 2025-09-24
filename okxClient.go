package galatvtr

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	// OKXBaseURL OKX API 基础 URL
	okxBaseUrl = "https://www.okx.com"
)

// OKXClient OKX API 客户端
type OKXClient struct {
	Client     *http.Client
	BaseUrl    string
	apiKey     string
	apiSecret  string
	passphrase string
	isTestnet  int
}

func NewOkClientWithoutKey(baseUrl string) *OKXClient {
	if baseUrl == "" {
		baseUrl = okxBaseUrl
	}
	return &OKXClient{
		Client: &http.Client{
			Timeout: time.Second * 10,
		},
		BaseUrl: baseUrl,
	}
}

// NewOKXClient 创建一个新的 OKX API 客户端
func NewOKXClient(baseUrl, apiKey, apiSecret, passphrase string, isTestnet int) *OKXClient {
	if baseUrl == "" {
		baseUrl = okxBaseUrl
	}
	return &OKXClient{
		Client: &http.Client{
			Timeout: time.Second * 10,
		},
		BaseUrl:    baseUrl,
		apiKey:     apiKey,
		apiSecret:  apiSecret,
		passphrase: passphrase,
		isTestnet:  isTestnet,
	}
}

// 发送请求到 OKX API
func (c *OKXClient) SendRequest(method, endpoint string, params interface{}) ([]byte, error) {
	var reqBody []byte
	var err error

	url := c.BaseUrl + endpoint

	if params != nil && (method == "POST" || method == "PUT") {
		reqBody, err = json.Marshal(params)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	// preHash := timestamp + method + endpoint
	signature := c.sign(c.apiSecret, timestamp, method, endpoint, reqBody)

	req.Header.Set("OK-ACCESS-KEY", c.apiKey)
	req.Header.Set("OK-ACCESS-SIGN", signature)
	req.Header.Set("OK-ACCESS-TIMESTAMP", timestamp)
	req.Header.Set("OK-ACCESS-PASSPHRASE", c.passphrase)
	req.Header.Set("Content-Type", "application/json")

	if c.isTestnet == 1 {
		req.Header.Set("x-simulated-trading", "1")
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API 请求失败: %s, 状态码: %d", string(body), resp.StatusCode)
	}

	return body, nil
}

func (c *OKXClient) SendRequestNoAuth(method, endpoint string, params interface{}) ([]byte, int, error) {
	var reqBody []byte
	var err error

	url := c.BaseUrl + endpoint

	if params != nil && (method == "POST" || method == "PUT") {
		reqBody, err = json.Marshal(params)
		if err != nil {
			return nil, -1, err
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, -1, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, -1, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, -1, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, resp.StatusCode, fmt.Errorf("API 请求失败: %s, 状态码: %d", string(body), resp.StatusCode)
	}

	return body, resp.StatusCode, nil
}

// 生成 OKX API 请求所需的签名
func (c *OKXClient) sign(secretKey, timestamp, method, requestPath string, body []byte) string {
	message := timestamp + method + requestPath
	if method == "POST" || method == "PUT" {
		message += string(body)
	}

	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
