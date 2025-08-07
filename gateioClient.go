package galatvtr

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/antihax/optional"
	"github.com/gateio/gateapi-go/v6"
)

const (
	GateIOBaseURL    = "https://api.gateio.ws/api/v4"
	GateIOTestnetURL = "https://api-testnet.gateapi.io/api/v4"
)

type GateIOClient struct {
	Client *gateapi.APIClient
	Ctx    context.Context
}

func NewGateIOClient(apiKey, secretKey string, isTestnet bool) *GateIOClient {
	configuration := gateapi.NewConfiguration()
	// 设置HTTP客户端超时
	configuration.HTTPClient = &http.Client{
		Timeout: time.Second * 30,
	}

	// 创建API客户端
	client := gateapi.NewAPIClient(configuration)

	// 设置API基础URL
	if isTestnet {
		client.ChangeBasePath(GateIOTestnetURL)
	}

	// 创建认证上下文
	ctx := context.WithValue(context.Background(), gateapi.ContextGateAPIV4, gateapi.GateAPIV4{
		Key:    apiKey,
		Secret: secretKey,
	})

	return &GateIOClient{
		Client: client,
		Ctx:    ctx,
	}
}

func NewGateIOClientWithoutAuth() *GateIOClient {

	// 创建API客户端
	client := gateapi.NewAPIClient(gateapi.NewConfiguration())

	return &GateIOClient{
		Client: client,
		Ctx:    context.Background(),
	}
}

// GetTicker 获取单个产品行情信息
func (g *GateIOClient) GetTicker(currencyPair string) (*gateapi.Ticker, error) {
	tickers, _, err := g.Client.SpotApi.ListTickers(g.Ctx, &gateapi.ListTickersOpts{
		CurrencyPair: optional.NewString(currencyPair),
	})
	if err != nil {
		return nil, err
	}
	if len(tickers) == 0 {
		return nil, fmt.Errorf("ticker not found for currency pair: %s", currencyPair)
	}
	return &tickers[0], nil
}

func (g *GateIOClient) GetTickerLast(currencyPair string) (string, error) {

	instId, errInstId := convertTradingViewTickerToGateioInstId(currencyPair)
	if errInstId != nil {
		return "", errInstId
	}

	tickers, _, err := g.Client.SpotApi.ListTickers(g.Ctx, &gateapi.ListTickersOpts{
		CurrencyPair: optional.NewString(instId),
	})
	if err != nil {
		return "", err
	}
	if len(tickers) == 0 {
		return "", fmt.Errorf("ticker not found for currency pair: %s", instId)
	}
	return tickers[0].Last, nil
}

// GetBalance 获取账户余额
func (g *GateIOClient) GetBalance() ([]gateapi.SpotAccount, error) {
	accounts, _, err := g.Client.SpotApi.ListSpotAccounts(g.Ctx, &gateapi.ListSpotAccountsOpts{})
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

// PlaceOrder 下单
func (g *GateIOClient) PlaceSpotOrder(order gateapi.Order) (*gateapi.Order, error) {
	result, _, err := g.Client.SpotApi.CreateOrder(g.Ctx, order, &gateapi.CreateOrderOpts{})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// PlaceOrder 下单
func (g *GateIOClient) PlaceFutureOrder(order gateapi.FuturesOrder) (*gateapi.FuturesOrder, *http.Response, error) {
	result, httpRes, err := g.Client.FuturesApi.CreateFuturesOrder(g.Ctx, "usdt", order, nil)
	if err != nil {
		return nil, httpRes, err
	}
	return &result, httpRes, nil
}

// GetOrderStatus 查询订单状态
func (g *GateIOClient) GetOrderStatus(orderId, currencyPair string) (*gateapi.Order, error) {
	order, _, err := g.Client.SpotApi.GetOrder(g.Ctx, orderId, currencyPair, &gateapi.GetOrderOpts{})
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// CancelOrder 撤销订单
func (g *GateIOClient) CancelOrder(orderId, currencyPair string) (*gateapi.Order, error) {
	order, _, err := g.Client.SpotApi.CancelOrder(g.Ctx, orderId, currencyPair, &gateapi.CancelOrderOpts{})
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (g *GateIOClient) GetPosition(instId string) (string, error) {
	position, _, err := g.Client.FuturesApi.GetPosition(g.Ctx, "usdt", instId)
	if err != nil {
		return "", err
	}
	return position.Leverage, nil
}

func (g *GateIOClient) SetPositionLever(instId, lever string) error {
	_, _, err := g.Client.FuturesApi.UpdatePositionLeverage(g.Ctx, "usdt", instId, lever, nil)
	if err != nil {
		return err
	}
	return nil
}

func (g *GateIOClient) GetInstruments(instId string) (float64, error) {
	res, _, err := g.Client.FuturesApi.GetFuturesContract(g.Ctx, "usdt", instId)
	if err != nil {
		return 0, err
	}
	if valueFloat, err := strconv.ParseFloat(res.QuantoMultiplier, 64); err == nil {
		return valueFloat, nil
	} else {
		return 0, fmt.Errorf("解析value失败: %v", err)
	}
}
