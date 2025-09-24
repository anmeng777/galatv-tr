package galatvtr

import (
	"encoding/json"
	"fmt"
)

// GetAssetBalance 获取资金账户余额
func (c *OKXClient) GetAssetBalance(ccy string) (*AssetBalanceResponse, error) {
	endpoint := "/api/v5/asset/balances"

	// 构建查询参数
	if ccy != "" {
		endpoint += "?ccy=" + ccy
	}

	resp, err := c.SendRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result AssetBalanceResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if result.Code != "0" {
		return &result, fmt.Errorf("获取资金账户余额失败: %s", result.Msg)
	}

	return &result, nil
}

// AssetTransfer 资金划转
func (c *OKXClient) AssetTransfer(request AssetTransferRequest) (*AssetTransferResponse, error) {
	endpoint := "/api/v5/asset/transfer"

	// 打印请求参数
	fmt.Printf("资金划转参数: %+v\n", request)

	resp, err := c.SendRequest("POST", endpoint, request)
	if err != nil {
		return nil, err
	}

	var result AssetTransferResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if result.Code != "0" {
		return &result, fmt.Errorf("资金划转失败: %s", result.Msg)
	}

	return &result, nil
}
