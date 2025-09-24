package galatvtr

import (
	"encoding/json"
	"fmt"
)

// GetSavingsBalance 获取余币宝余额
func (c *OKXClient) GetSavingsBalance(ccy string) (*SavingsBalanceResponse, error) {
	endpoint := "/api/v5/finance/savings/balance"

	// 构建查询参数
	if ccy != "" {
		endpoint += "?ccy=" + ccy
	}

	resp, err := c.SendRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result SavingsBalanceResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if result.Code != "0" {
		return &result, fmt.Errorf("获取余币宝余额失败: %s", result.Msg)
	}

	return &result, nil
}

// SavingsPurchaseRedempt 余币宝申购/赎回
func (c *OKXClient) SavingsPurchaseRedempt(request SavingsPurchaseRedemptRequest) (*SavingsPurchaseRedemptResponse, error) {
	endpoint := "/api/v5/finance/savings/purchase-redempt"

	// 打印请求参数
	fmt.Printf("余币宝申购/赎回参数: %+v\n", request)

	resp, err := c.SendRequest("POST", endpoint, request)
	if err != nil {
		return nil, err
	}

	var result SavingsPurchaseRedemptResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if result.Code != "0" {
		return &result, fmt.Errorf("余币宝申购/赎回失败: %s", result.Msg)
	}

	return &result, nil
}
