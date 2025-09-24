package galatvtr

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// GetTicker 获取单个产品行情信息
func (c *OKXClient) GetTickerLast(instId string) (string, error) {
	endpoint := "/api/v5/market/ticker?instId=" + instId
	resp, err := c.SendRequest("GET", endpoint, nil)
	if err != nil {
		return "", err
	}
	var result TickerResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return "", err
	}
	if result.Code != "0" {
		return "", fmt.Errorf("获取行情信息失败: %s", result.Msg)
	}
	if len(result.Data) == 0 {
		return "", fmt.Errorf("未获取到行情数据")
	}
	return result.Data[0].Last, nil
}

func (c *OKXClient) OkGetKlineFecher(symbol, interval string, startTime, endTime *int64) ([][]string, error) {
	if interval == "6H" || interval == "12H" || interval == "1D" || interval == "2D" || interval == "3D" || interval == "1W" || interval == "1M" || interval == "3M" {
		interval = interval + "utc"
	}

	// 检查指针是否为nil，避免解引用nil指针
	startTimeStr := "nil"
	endTimeStr := "nil"
	if startTime != nil {
		startTimeStr = strconv.FormatInt(*startTime, 10)
	}
	if endTime != nil {
		endTimeStr = strconv.FormatInt(*endTime, 10)
	}
	fmt.Printf("------准备获取【%s】的K线数据%s~%s\n", interval, startTimeStr, endTimeStr)

	endpoint := "/api/v5/market/history-candles"

	params := "?instId=" + symbol + "&bar=" + interval
	if startTime != nil {
		params += "&before=" + strconv.FormatInt(*startTime, 10)
	}
	if endTime != nil {
		params += "&after=" + strconv.FormatInt(*endTime, 10)
	}
	params += "&limit=100"
	endpoint += params

	var body []byte
	for {
		var statusCode int
		var err2 error
		// 发送请求
		body, statusCode, err2 = c.SendRequestNoAuth("GET", endpoint, nil)
		if err2 != nil {
			fmt.Printf("发送请求错误: %v\n", err2)
			// return nil, -1, err2
			time.Sleep(1 * time.Second)
			continue
		}

		if statusCode == 429 {
			fmt.Printf("遇到接口限流(429)，等待后重试...\n")
			time.Sleep(1 * time.Second)
			continue
		} else {
			if statusCode != 200 {
				fmt.Printf("!!!!!!请求失败，间隔【%s】状态码: %d\n", interval, statusCode)
				// return nil, resp.StatusCode, err2
				time.Sleep(1 * time.Second)
				continue
			} else {
				break
			}
		}
	}

	var response HttpKlineResponse
	err := json.Unmarshal(body, &response)
	if err != nil {
		fmt.Printf("解析JSON错误: %v\n", err)
		return nil, err
	}

	if response.Code != "0" {
		fmt.Printf("API错误: %s\n", response.Msg)
		return nil, err
	}

	return response.Data, nil
}
