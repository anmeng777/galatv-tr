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
	"strconv"
	"time"
)

const (
	// OKXBaseURL OKX API 基础 URL
	okxBaseUrl = "https://www.okx.com"
)

// OKXClient OKX API 客户端
type OKXClient struct {
	Client  *http.Client
	BaseUrl string
}

// NewOKXClient 创建一个新的 OKX API 客户端
func NewOKXClient(baseUrl string) *OKXClient {
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

// 发送请求到 OKX API
func (c *OKXClient) SendRequest(apiKey, secretKey, passphrase string, isTestnet int, method, endpoint string, params interface{}) ([]byte, error) {
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
	signature := c.sign(secretKey, timestamp, method, endpoint, reqBody)

	req.Header.Set("OK-ACCESS-KEY", apiKey)
	req.Header.Set("OK-ACCESS-SIGN", signature)
	req.Header.Set("OK-ACCESS-TIMESTAMP", timestamp)
	req.Header.Set("OK-ACCESS-PASSPHRASE", passphrase)
	req.Header.Set("Content-Type", "application/json")

	if isTestnet == 1 {
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

	fmt.Println("API 响应: ", string(body))

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

	fmt.Println("API 响应: ", string(body))

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

// PlaceOrder 下单
func (c *OKXClient) PlaceOrder(apiKey, secretKey, passphrase string, isTestnet int, order OrderRequestOkx) (*OrderResponse, error) {
	endpoint := "/api/v5/trade/order"

	// 打印order
	fmt.Printf("下单参数: %+v\n", order)

	resp, err := c.SendRequest(apiKey, secretKey, passphrase, isTestnet, "POST", endpoint, order)
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

func (c *OKXClient) PlaceOrderHeyueOkx(apiKey, secretKey, passphrase string, isTestnet int, order OrderRequestOkx) (*OrderResponse, error) {
	endpoint := "/api/v5/trade/order"

	// 打印order
	fmt.Printf("下单参数: %+v\n", order)

	resp, err := c.SendRequest(apiKey, secretKey, passphrase, isTestnet, "POST", endpoint, order)
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

// GetTicker 获取单个产品行情信息
func (c *OKXClient) GetTickerLast(apiKey, secretKey, passphrase string, isTestnet int, instId string) (string, error) {
	endpoint := "/api/v5/market/ticker?instId=" + instId
	resp, err := c.SendRequest(apiKey, secretKey, passphrase, isTestnet, "GET", endpoint, nil)
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

// 查询指定币种余额
func (c *OKXClient) GetAccountBalance(apiKey, secretKey, passphrase string, isTestnet int, ccy string) (*BalanceResponse, error) {
	endpoint := "/api/v5/account/balance?ccy=" + ccy

	resp, err := c.SendRequest(apiKey, secretKey, passphrase, isTestnet, "GET", endpoint, nil)
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

// GetLeverageInfo 查询杠杆倍率信息
func (c *OKXClient) GetLeverageInfo(apiKey, secretKey, passphrase string, isTestnet int, instId, mgnMode string) (*LeverageInfoResponse, error) {
	endpoint := fmt.Sprintf("/api/v5/account/leverage-info?instId=%s&mgnMode=%s", instId, mgnMode)

	resp, err := c.SendRequest(apiKey, secretKey, passphrase, isTestnet, "GET", endpoint, nil)
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

	resp, err := c.SendRequest(apiKey, secretKey, passphrase, isTestnet, "POST", endpoint, request)
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

// GetOrderInfo 查询订单信息
func (c *OKXClient) GetOrderInfo(apiKey, secretKey, passphrase string, isTestnet int, instId, ordId, clOrdId string) (*OrderInfoResponse, error) {
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

	resp, err := c.SendRequest(apiKey, secretKey, passphrase, isTestnet, "GET", endpoint, nil)
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
func (c *OKXClient) PlaceAlgoOrder(apiKey, secretKey, passphrase string, isTestnet int, order AlgoOrderRequest) (*AlgoOrderResponse, error) {
	endpoint := "/api/v5/trade/order-algo"

	// 打印策略委托下单参数
	fmt.Printf("策略委托下单参数: %+v\n", order)

	resp, err := c.SendRequest(apiKey, secretKey, passphrase, isTestnet, "POST", endpoint, order)
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

// GetPositions 获取持仓信息
func (c *OKXClient) GetPositions(apiKey, secretKey, passphrase string, isTestnet int, instType, instId string) (*PositionsResponse, error) {
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

	resp, err := c.SendRequest(apiKey, secretKey, passphrase, isTestnet, "GET", endpoint, nil)
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
func (c *OKXClient) GetAlgoOrdersPending(apiKey, secretKey, passphrase string, isTestnet int, ordType, instType, instId string) (*AlgoOrdersPendingResponse, error) {
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

	resp, err := c.SendRequest(apiKey, secretKey, passphrase, isTestnet, "GET", endpoint, nil)
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
func (c *OKXClient) CancelAlgoOrders(apiKey, secretKey, passphrase string, isTestnet int, requests []CancelAlgoOrderRequest) (*CancelAlgoOrdersResponse, error) {
	endpoint := "/api/v5/trade/cancel-algos"

	// 打印撤销策略委托订单参数
	fmt.Printf("撤销策略委托订单参数: %+v\n", requests)

	resp, err := c.SendRequest(apiKey, secretKey, passphrase, isTestnet, "POST", endpoint, requests)
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

// GetSavingsBalance 获取余币宝余额
func (c *OKXClient) GetSavingsBalance(apiKey, secretKey, passphrase string, isTestnet int, ccy string) (*SavingsBalanceResponse, error) {
	endpoint := "/api/v5/finance/savings/balance"

	// 构建查询参数
	if ccy != "" {
		endpoint += "?ccy=" + ccy
	}

	resp, err := c.SendRequest(apiKey, secretKey, passphrase, isTestnet, "GET", endpoint, nil)
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
func (c *OKXClient) SavingsPurchaseRedempt(apiKey, secretKey, passphrase string, isTestnet int, request SavingsPurchaseRedemptRequest) (*SavingsPurchaseRedemptResponse, error) {
	endpoint := "/api/v5/finance/savings/purchase-redempt"

	// 打印请求参数
	fmt.Printf("余币宝申购/赎回参数: %+v\n", request)

	resp, err := c.SendRequest(apiKey, secretKey, passphrase, isTestnet, "POST", endpoint, request)
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

// GetAssetBalance 获取资金账户余额
func (c *OKXClient) GetAssetBalance(apiKey, secretKey, passphrase string, isTestnet int, ccy string) (*AssetBalanceResponse, error) {
	endpoint := "/api/v5/asset/balances"

	// 构建查询参数
	if ccy != "" {
		endpoint += "?ccy=" + ccy
	}

	resp, err := c.SendRequest(apiKey, secretKey, passphrase, isTestnet, "GET", endpoint, nil)
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
func (c *OKXClient) AssetTransfer(apiKey, secretKey, passphrase string, isTestnet int, request AssetTransferRequest) (*AssetTransferResponse, error) {
	endpoint := "/api/v5/asset/transfer"

	// 打印请求参数
	fmt.Printf("资金划转参数: %+v\n", request)

	resp, err := c.SendRequest(apiKey, secretKey, passphrase, isTestnet, "POST", endpoint, request)
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

func ZhuanbiRedemptionAllToAccountBalance(client *OKXClient, apiKey, secretKey, passphrase string, isTestnet int, ticker string) error {
	instId, err := convertTvTrickerToSingleCoinName(ticker)
	if err != nil {
		fmt.Printf("[Redemption] 转换交易对失败: %v\n", err)
		return err
	}

	var assetBalance float64
	getAssetBalanceResult, err := client.GetAssetBalance(apiKey, secretKey, passphrase, isTestnet, instId)
	if err != nil {
		fmt.Printf("[Redemption] 查询资金账户余额失败: %v\n", err)
		return err
	}
	for _, data := range getAssetBalanceResult.Data {
		if data.Ccy == instId {
			if valueFloat, err := strconv.ParseFloat(data.AvailBal, 64); err == nil {
				assetBalance = valueFloat
			}
		}
	}

	getSavingsBalanceResult, err := client.GetSavingsBalance(apiKey, secretKey, passphrase, isTestnet, instId)
	if err != nil {
		fmt.Printf("[Redemption] 查询稳定赚币余额失败: %v\n", err)
		return err
	}
	var savingBalance float64
	for _, data := range getSavingsBalanceResult.Data {
		if data.Ccy == instId {
			if valueFloat, err := strconv.ParseFloat(data.Amt, 64); err == nil {
				savingBalance = valueFloat
			}
		}
	}

	if savingBalance > 0 {
		request := SavingsPurchaseRedemptRequest{
			Ccy:  instId,
			Side: "redempt",
			Amt:  strconv.FormatFloat(zhuanbiFormatFloat(instId, savingBalance), 'f', 8, 64),
		}
		// 赎回到资金账户
		fmt.Printf("[Redemption] 开始从稳定赚币赎回...\n")
		_, errS := client.SavingsPurchaseRedempt(apiKey, secretKey, passphrase, isTestnet, request)
		if errS == nil {
			// 等待赎回成功，一直查询资金账户直到余额变化
			fmt.Printf("[Redemption] 等待赎回到账，监控资金账户余额变化...\n")
			maxRetries := 4 // 最多等待60次，每次间隔500毫秒，总共30秒
			for i := 0; i < maxRetries; i++ {
				time.Sleep(500 * time.Millisecond) // 等待500毫秒

				// 查询当前资金账户余额
				currentAssetBalanceResult, err := client.GetAssetBalance(apiKey, secretKey, passphrase, isTestnet, instId)
				if err != nil {
					fmt.Printf("[Redemption] 第%d次查询资金账户余额失败: %v\n", i+1, err)
					continue
				}

				var currentAssetBalance float64
				for _, data := range currentAssetBalanceResult.Data {
					if data.Ccy == instId {
						if valueFloat, err := strconv.ParseFloat(data.AvailBal, 64); err == nil {
							currentAssetBalance = valueFloat
						}
					}
				}

				fmt.Printf("[Redemption] 第%d次查询，当前资金账户余额: %.8f，原余额: %.8f\n", i+1, currentAssetBalance, assetBalance)

				// 如果余额有增加，说明赎回已到账
				if currentAssetBalance > assetBalance {
					fmt.Printf("[Redemption] 检测到余额增加，赎回已到账\n")
					assetBalance = currentAssetBalance
					break
				}
			}
		}
	}

	transferRequest := AssetTransferRequest{
		Ccy:  instId,
		Amt:  strconv.FormatFloat(zhuanbiFormatFloat(instId, assetBalance), 'f', 8, 64),
		From: "6",  // 资金账户
		To:   "18", // 交易账户
	}
	_, errT := client.AssetTransfer(apiKey, secretKey, passphrase, isTestnet, transferRequest)
	if errT != nil {
		fmt.Printf("[Redemption] 资金划转失败: %v\n", errT)
		return errT
	}

	return nil
}
