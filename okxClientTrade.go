package galatvtr

import (
	"encoding/json"
	"fmt"
)

// PlaceOrder 下单
func (c *OKXClient) PlaceOrder(order OrderRequestOkx) (*OrderResponse, error) {
	endpoint := "/api/v5/trade/order"

	// 打印order
	fmt.Printf("下单参数: %+v\n", order)

	resp, err := c.SendRequest("POST", endpoint, order)
	if err != nil {
		return nil, err
	}

	var result OrderResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if result.Code != "0" {
		return &result, fmt.Errorf("下单失败: %s", result.Msg)
	}

	return &result, nil
}

// 合约下单
func (c *OKXClient) PlaceOrderHeyueOkx(order OrderRequestOkx) (*OrderResponse, error) {
	endpoint := "/api/v5/trade/order"

	// 打印order
	fmt.Printf("下单参数: %+v\n", order)

	resp, err := c.SendRequest("POST", endpoint, order)
	if err != nil {
		return nil, err
	}

	var result OrderResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if result.Code != "0" {
		return &result, fmt.Errorf("下单失败: %s", result.Msg)
	}

	return &result, nil
}

// GetOrderInfo 查询订单信息
func (c *OKXClient) GetOrderInfo(instId, ordId, clOrdId string) (*OrderInfoResponse, error) {
	endpoint := "/api/v5/trade/order"

	// 构建查询参数
	if instId == "" {
		return nil, fmt.Errorf("instId 参数不能为空")
	}

	if ordId != "" {
		endpoint += "?instId=" + instId + "&ordId=" + ordId
	} else if clOrdId != "" {
		endpoint += "?instId=" + instId + "&clOrdId=" + clOrdId
	} else {
		return nil, fmt.Errorf("ordId 和 clOrdId 至少需要提供一个")
	}

	resp, err := c.SendRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result OrderInfoResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if result.Code != "0" {
		return &result, fmt.Errorf("查询订单信息失败: %s", result.Msg)
	}

	return &result, nil
}

// PlaceAlgoOrder 策略委托下单
func (c *OKXClient) PlaceAlgoOrder(order AlgoOrderRequest) (*AlgoOrderResponse, error) {
	endpoint := "/api/v5/trade/order-algo"

	// 打印策略委托下单参数
	fmt.Printf("策略委托下单参数: %+v\n", order)

	resp, err := c.SendRequest("POST", endpoint, order)
	if err != nil {
		return nil, err
	}

	var result AlgoOrderResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if result.Code != "0" {
		return &result, fmt.Errorf("策略委托下单失败: %s", result.Msg)
	}

	return &result, nil
}

// GetAlgoOrderInfo 查询策略委托单信息
// func (c *OKXClient) GetAlgoOrderInfo(apiKey, secretKey, passphrase string, isTestnet int, algoId, algoClOrdId string) (*galastruct.AlgoOrderInfoResponse, error) {
// 	endpoint := "/api/v5/trade/order-algo"

// 	// 构建查询参数
// 	if algoId != "" {
// 		endpoint += "?algoId=" + algoId
// 	} else if algoClOrdId != "" {
// 		endpoint += "?algoClOrdId=" + algoClOrdId
// 	} else {
// 		return nil, fmt.Errorf("algoId 和 algoClOrdId 至少需要提供一个")
// 	}

// 	resp, err := c.SendRequest(apiKey, secretKey, passphrase, isTestnet, "GET", endpoint, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var result galastruct.AlgoOrderInfoResponse
// 	if err := json.Unmarshal(resp, &result); err != nil {
// 		return nil, err
// 	}

// 	if result.Code != "0" {
// 		if result.Msg == "Order does not exist" {
// 			return nil, nil
// 		}
// 		return &result, fmt.Errorf("查询策略委托单信息失败: %s", result.Msg)
// 	}

// 	return &result, nil
// }

// GetAlgoOrdersPending 获取未完成策略委托单列表
func (c *OKXClient) GetAlgoOrdersPending(ordType, instType, instId string) (*AlgoOrdersPendingResponse, error) {
	endpoint := "/api/v5/trade/orders-algo-pending"

	// 构建查询参数
	params := "?ordType=" + ordType
	if instType != "" {
		params += "&instType=" + instType
	}
	if instId != "" {
		params += "&instId=" + instId
	}
	endpoint += params

	resp, err := c.SendRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result AlgoOrdersPendingResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if result.Code != "0" {
		return &result, fmt.Errorf("获取未完成策略委托单列表失败: %s", result.Msg)
	}

	return &result, nil
}

// CancelAlgoOrders 撤销策略委托订单
func (c *OKXClient) CancelAlgoOrders(requests []CancelAlgoOrderRequest) (*CancelAlgoOrdersResponse, error) {
	endpoint := "/api/v5/trade/cancel-algos"

	// 打印撤销策略委托订单参数
	fmt.Printf("撤销策略委托订单参数: %+v\n", requests)

	resp, err := c.SendRequest("POST", endpoint, requests)
	if err != nil {
		return nil, err
	}

	var result CancelAlgoOrdersResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if result.Code != "0" {
		return &result, fmt.Errorf("撤销策略委托订单失败: %s", result.Msg)
	}

	return &result, nil
}

// GetOrdersHistoryArchive 获取历史订单记录（近三个月）
func (c *OKXClient) GetOrdersHistoryArchive(request OrdersHistoryArchiveRequest) (*OrdersHistoryArchiveResponse, error) {
	// 构建查询参数
	queryParams := ""
	if request.InstType != "" {
		queryParams += "instType=" + request.InstType
	}
	if request.InstFamily != "" {
		if queryParams != "" {
			queryParams += "&"
		}
		queryParams += "instFamily=" + request.InstFamily
	}
	if request.InstId != "" {
		if queryParams != "" {
			queryParams += "&"
		}
		queryParams += "instId=" + request.InstId
	}
	if request.OrdType != "" {
		if queryParams != "" {
			queryParams += "&"
		}
		queryParams += "ordType=" + request.OrdType
	}
	if request.State != "" {
		if queryParams != "" {
			queryParams += "&"
		}
		queryParams += "state=" + request.State
	}
	if request.Category != "" {
		if queryParams != "" {
			queryParams += "&"
		}
		queryParams += "category=" + request.Category
	}
	if request.After != "" {
		if queryParams != "" {
			queryParams += "&"
		}
		queryParams += "after=" + request.After
	}
	if request.Before != "" {
		if queryParams != "" {
			queryParams += "&"
		}
		queryParams += "before=" + request.Before
	}
	if request.Begin != "" {
		if queryParams != "" {
			queryParams += "&"
		}
		queryParams += "begin=" + request.Begin
	}
	if request.End != "" {
		if queryParams != "" {
			queryParams += "&"
		}
		queryParams += "end=" + request.End
	}
	if request.Limit != "" {
		if queryParams != "" {
			queryParams += "&"
		}
		queryParams += "limit=" + request.Limit
	}

	// 构建完整的endpoint
	endpoint := "/api/v5/trade/orders-history-archive"
	if queryParams != "" {
		endpoint += "?" + queryParams
	}

	// 发送GET请求
	body, err := c.SendRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	// 解析响应
	var response OrdersHistoryArchiveResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
