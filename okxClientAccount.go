package galatvtr

import (
	"encoding/json"
	"fmt"
)

// 查询指定币种余额
func (c *OKXClient) GetAccountBalance(ccy string) (*BalanceResponse, error) {
	endpoint := "/api/v5/account/balance?ccy=" + ccy

	resp, err := c.SendRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	var result BalanceResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if result.Code != "0" {
		return nil, fmt.Errorf("查询余额失败: %s", result.Msg)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("未获取到余额数据")
	}

	return &result, nil
}

// GetPositions 获取持仓信息
func (c *OKXClient) GetPositions(instType, instId string) (*PositionsResponse, error) {
	endpoint := "/api/v5/account/positions"

	// 构建查询参数
	if instType != "" {
		endpoint += "?instType=" + instType
		if instId != "" {
			endpoint += "&instId=" + instId
		}
	} else if instId != "" {
		endpoint += "?instId=" + instId
	}

	resp, err := c.SendRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result PositionsResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if result.Code != "0" {
		return &result, fmt.Errorf("获取持仓信息失败: %s", result.Msg)
	}

	return &result, nil
}

// GetLeverageInfo 查询杠杆倍率信息
func (c *OKXClient) GetLeverageInfo(instId, mgnMode string) (*LeverageInfoResponse, error) {
	endpoint := fmt.Sprintf("/api/v5/account/leverage-info?instId=%s&mgnMode=%s", instId, mgnMode)

	resp, err := c.SendRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result LeverageInfoResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if result.Code != "0" {
		return &result, fmt.Errorf("查询杠杆倍率失败: %s", result.Msg)
	}

	return &result, nil
}

// SetLeverage 设置杠杆倍率
func (c *OKXClient) SetLeverage(apiKey, secretKey, passphrase string, isTestnet int, request SetLeverageRequest) (*SetLeverageResponse, error) {
	endpoint := "/api/v5/account/set-leverage"

	resp, err := c.SendRequest("POST", endpoint, request)
	if err != nil {
		return nil, err
	}

	var result SetLeverageResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if result.Code != "0" {
		return &result, fmt.Errorf("设置杠杆倍率失败: %s", result.Msg)
	}

	return &result, nil
}
