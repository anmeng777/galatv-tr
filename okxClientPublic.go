package galatvtr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// GetInstruments 获取交易产品基础信息
func (c *OKXClient) GetInstruments(instType, instId string) (float64, error) {
	endpoint := "/api/v5/public/instruments"

	// 构建查询参数
	if instType == "" {
		return 0, fmt.Errorf("instType 参数不能为空")
	}

	params := "?instType=" + instType
	if instId != "" {
		params += "&instId=" + instId
	}
	endpoint += params

	// 使用无认证请求，因为这是公共接口
	resp, statusCode, err := c.SendRequestNoAuth("GET", endpoint, nil)
	if err != nil || statusCode != http.StatusOK {
		return 0, err
	}

	var result InstrumentsResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if result.Code != "0" {
		return 0, fmt.Errorf("获取交易产品基础信息失败: %s", result.Msg)
	}
	if len(result.Data) > 0 {
		if valueFloat, err := strconv.ParseFloat(result.Data[0].CtVal, 64); err == nil {
			return valueFloat, nil
		} else {
			return 0, fmt.Errorf("解析value失败: %v", err)
		}
	} else {
		return 0, fmt.Errorf("获取交易产品基础信息失败: %s", result.Msg)
	}
}
