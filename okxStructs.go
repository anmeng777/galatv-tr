package galatvtr

// OrderRequestOkx 下单请求参数
type OrderRequestOkx struct {
	InstID         string          `json:"instId"`                   // 产品ID
	TdMode         string          `json:"tdMode"`                   // 交易模式
	Side           string          `json:"side"`                     // 订单方向
	PosSide        string          `json:"posSide,omitempty"`        // 持仓方向，永续合约必填
	OrdType        string          `json:"ordType"`                  // 订单类型
	Sz             string          `json:"sz"`                       // 委托数量
	Px             string          `json:"px,omitempty"`             // 委托价格，仅适用于限价单
	TgtCcy         string          `json:"tgtCcy,omitempty"`         // 计价货币
	ReduceOnly     bool            `json:"reduceOnly,omitempty"`     // 是否只减仓
	Lever          string          `json:"lever,omitempty"`          // 杠杆倍率
	ClOrdID        string          `json:"clOrdId,omitempty"`        // 客户自定义订单ID
	AttachAlgoOrds []AttachAlgoOrd `json:"attachAlgoOrds,omitempty"` // 下单附带止盈止损信息
}

// AttachAlgoOrd 下单附带止盈止损信息
type AttachAlgoOrd struct {
	// AttachAlgoClOrdId string `json:"attachAlgoClOrdId,omitempty"` // 客户自定义的策略订单ID
	TpTriggerPx string `json:"tpTriggerPx,omitempty"` // 止盈触发价
	TpOrdPx     string `json:"tpOrdPx,omitempty"`     // 止盈委托价
	// TpOrdKind         string `json:"tpOrdKind,omitempty"`         // 止盈订单类型
	// SlTriggerPx       string `json:"slTriggerPx,omitempty"`       // 止损触发价
	// SlOrdPx           string `json:"slOrdPx,omitempty"`           // 止损委托价
	// TpTriggerPxType   string `json:"tpTriggerPxType,omitempty"`   // 止盈触发价类型
	// SlTriggerPxType   string `json:"slTriggerPxType,omitempty"`   // 止损触发价类型
	// Sz                string `json:"sz,omitempty"`                // 数量
}

// OrderResponse 下单响应
type OrderResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		ClOrdID string `json:"clOrdId"`
		OrdID   string `json:"ordId"`
		Tag     string `json:"tag"`
		SCode   string `json:"sCode"`
		SMsg    string `json:"sMsg"`
	} `json:"data"`
}

// TickerData 行情数据
type TickerData struct {
	InstType  string `json:"instType"`  // 产品类型
	InstId    string `json:"instId"`    // 产品ID
	Last      string `json:"last"`      // 最新成交价
	LastSz    string `json:"lastSz"`    // 最新成交的数量
	AskPx     string `json:"askPx"`     // 卖一价
	AskSz     string `json:"askSz"`     // 卖一价对应的数量
	BidPx     string `json:"bidPx"`     // 买一价
	BidSz     string `json:"bidSz"`     // 买一价对应的数量
	Open24h   string `json:"open24h"`   // 24小时开盘价
	High24h   string `json:"high24h"`   // 24小时最高价
	Low24h    string `json:"low24h"`    // 24小时最低价
	VolCcy24h string `json:"volCcy24h"` // 24小时成交量，以币为单位
	Vol24h    string `json:"vol24h"`    // 24小时成交量，以张为单位
	SodUtc0   string `json:"sodUtc0"`   // UTC+0 时开盘价
	SodUtc8   string `json:"sodUtc8"`   // UTC+8 时开盘价
	Ts        string `json:"ts"`        // ticker数据产生时间
}

// TickerResponse 行情信息响应
type TickerResponse struct {
	Code string       `json:"code"`
	Msg  string       `json:"msg"`
	Data []TickerData `json:"data"`
}

// 余额返回 BalanceResponse
type BalanceResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Details []struct {
			Ccy       string `json:"ccy"`
			AvailEq   string `json:"availEq"`
			AvailBal  string `json:"availBal"`
			CashBal   string `json:"cashBal"`
			DisEq     string `json:"disEq"`
			Eq        string `json:"eq"`
			EqUsd     string `json:"eqUsd"`
			FrozenBal string `json:"frozenBal"`
		}
	}
}

// 杠杆倍率查询响应
type LeverageInfoResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Ccy     string `json:"ccy"`
		InstId  string `json:"instId"`
		MgnMode string `json:"mgnMode"`
		PosSide string `json:"posSide"`
		Lever   string `json:"lever"`
	} `json:"data"`
}

// 设置杠杆倍率响应
type SetLeverageResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Lever   string `json:"lever"`
		MgnMode string `json:"mgnMode"`
		InstId  string `json:"instId"`
		PosSide string `json:"posSide"`
	} `json:"data"`
}

// 设置杠杆倍率请求
type SetLeverageRequest struct {
	InstId  string `json:"instId,omitempty"`  // 产品ID
	Ccy     string `json:"ccy,omitempty"`     // 币种
	Lever   string `json:"lever"`             // 杠杆倍数
	MgnMode string `json:"mgnMode"`           // 保证金模式
	PosSide string `json:"posSide,omitempty"` // 持仓方向
}

// 策略委托下单请求
type AlgoOrderRequest struct {
	InstId        string `json:"instId"`                  // 产品ID
	TdMode        string `json:"tdMode"`                  // 交易模式
	Side          string `json:"side"`                    // 订单方向
	PosSide       string `json:"posSide,omitempty"`       // 持仓方向
	OrdType       string `json:"ordType"`                 // 订单类型
	Sz            string `json:"sz,omitempty"`            // 委托数量
	CloseFraction string `json:"closeFraction,omitempty"` // 平仓百分比
	Tag           string `json:"tag,omitempty"`           // 订单标签
	TgtCcy        string `json:"tgtCcy,omitempty"`        // 委托数量类型
	AlgoClOrdId   string `json:"algoClOrdId,omitempty"`   // 客户自定义策略订单ID
	ReduceOnly    bool   `json:"reduceOnly,omitempty"`    // 是否只减仓
	// 止盈止损参数
	TpTriggerPx     string `json:"tpTriggerPx,omitempty"`     // 止盈触发价
	TpTriggerPxType string `json:"tpTriggerPxType,omitempty"` // 止盈触发价类型
	TpOrdPx         string `json:"tpOrdPx,omitempty"`         // 止盈委托价
	TpOrdKind       string `json:"tpOrdKind,omitempty"`       // 止盈订单类型
	SlTriggerPx     string `json:"slTriggerPx,omitempty"`     // 止损触发价
	SlTriggerPxType string `json:"slTriggerPxType,omitempty"` // 止损触发价类型
	SlOrdPx         string `json:"slOrdPx,omitempty"`         // 止损委托价
	CxlOnClosePos   bool   `json:"cxlOnClosePos,omitempty"`   // 是否与仓位关联
}

// 策略委托下单响应
type AlgoOrderResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		AlgoId      string `json:"algoId"`
		ClOrdId     string `json:"clOrdId"`
		AlgoClOrdId string `json:"algoClOrdId"`
		SCode       string `json:"sCode"`
		SMsg        string `json:"sMsg"`
		Tag         string `json:"tag"`
	} `json:"data"`
}

// PositionData 单个持仓信息
type PositionData struct {
	InstType               string `json:"instType"`               // 产品类型
	InstId                 string `json:"instId"`                 // 产品ID
	Lever                  string `json:"lever"`                  // 杠杆倍率
	AvailPos               string `json:"availPos"`               // 可平仓数量
	PosSide                string `json:"posSide"`                // 持仓方向
	Pos                    string `json:"pos"`                    // 持仓数量
	BaseCcy                string `json:"baseCcy"`                // 交易货币币种
	QuoteCcy               string `json:"quoteCcy"`               // 计价货币币种
	PosCcy                 string `json:"posCcy"`                 // 持仓币种
	AvgPx                  string `json:"avgPx"`                  // 开仓均价
	MarkPx                 string `json:"markPx"`                 // 最新标记价格
	Upl                    string `json:"upl"`                    // 未实现收益
	UplRatio               string `json:"uplRatio"`               // 未实现收益率
	MgnMode                string `json:"mgnMode"`                // 保证金模式
	NotionalUsd            string `json:"notionalUsd"`            // 以美元价值为单位的持仓数量
	Adl                    string `json:"adl"`                    // 信号区
	Imr                    string `json:"imr"`                    // 初始保证金
	Mmr                    string `json:"mmr"`                    // 维持保证金
	MgnRatio               string `json:"mgnRatio"`               // 保证金率
	MgnMulti               string `json:"mgnMulti"`               // 保证金倍数
	Liab                   string `json:"liab"`                   // 负债额
	LiabCcy                string `json:"liabCcy"`                // 负债币种
	SzLmt                  string `json:"szLmt"`                  // 持仓限制
	UTime                  string `json:"uTime"`                  // 最近一次持仓更新时间
	CTime                  string `json:"cTime"`                  // 持仓创建时间
	PTime                  string `json:"pTime"`                  // 推送时间
	BizRefId               string `json:"bizRefId"`               // 外部业务id
	BizRefType             string `json:"bizRefType"`             // 外部业务类型
	LiqPx                  string `json:"liqPx"`                  // 强平价格
	TradeId                string `json:"tradeId"`                // 最新成交ID
	OptVal                 string `json:"optVal"`                 // 期权价值
	PendingCloseOrdLiabVal string `json:"pendingCloseOrdLiabVal"` // 挂单平仓负债价值
	SpotInUseAmt           string `json:"spotInUseAmt"`           // 现货对冲占用数量
	SpotInUseCcy           string `json:"spotInUseCcy"`           // 现货对冲占用币种
	ClSpotInUseAmt         string `json:"clSpotInUseAmt"`         // 用户自定义现货对冲占用数量
	MaxSpotInUseAmt        string `json:"maxSpotInUseAmt"`        // 最大可用现货对冲数量
	DeltaBS                string `json:"deltaBS"`                // delta值
	DeltaPA                string `json:"deltaPA"`                // delta值（按组合计算）
	GammaBS                string `json:"gammaBS"`                // gamma值
	GammaPA                string `json:"gammaPA"`                // gamma值（按组合计算）
	ThetaBS                string `json:"thetaBS"`                // theta值
	ThetaPA                string `json:"thetaPA"`                // theta值（按组合计算）
	VegaBS                 string `json:"vegaBS"`                 // vega值
	VegaPA                 string `json:"vegaPA"`                 // vega值（按组合计算）
	RealizedPnl            string `json:"realizedPnl"`            // 已实现收益
	Pnl                    string `json:"pnl"`                    // 收益
	Fee                    string `json:"fee"`                    // 累计手续费
	FundingFee             string `json:"fundingFee"`             // 累计资金费用
	LiqPenalty             string `json:"liqPenalty"`             // 累计爆仓罚金
	CloseOrderAlgo         []struct {
		AlgoId        string `json:"algoId"`        // 策略委托订单ID
		SlTriggerPx   string `json:"slTriggerPx"`   // 止损触发价
		SlOrdPx       string `json:"slOrdPx"`       // 止损委托价
		TpTriggerPx   string `json:"tpTriggerPx"`   // 止盈触发价
		TpOrdPx       string `json:"tpOrdPx"`       // 止盈委托价
		CloseFraction string `json:"closeFraction"` // 平仓分数
	} `json:"closeOrderAlgo"` // 平仓策略委托订单
}

// PositionsResponse 持仓信息响应
type PositionsResponse struct {
	Code string         `json:"code"`
	Msg  string         `json:"msg"`
	Data []PositionData `json:"data"`
}

// AlgoOrdersPendingData 未完成策略委托单数据
type AlgoOrdersPendingData struct {
	ActivePx             string `json:"activePx"`             // 激活价格
	ActualPx             string `json:"actualPx"`             // 实际委托价
	ActualSide           string `json:"actualSide"`           // 实际触发方向
	ActualSz             string `json:"actualSz"`             // 实际委托量
	AlgoClOrdId          string `json:"algoClOrdId"`          // 客户自定义策略订单ID
	AlgoId               string `json:"algoId"`               // 策略委托单ID
	AmendPxOnTriggerType string `json:"amendPxOnTriggerType"` // 触发时修改价格类型
	CTime                string `json:"cTime"`                // 订单创建时间
	InstId               string `json:"instId"`               // 产品ID
	InstType             string `json:"instType"`             // 产品类型
	Lever                string `json:"lever"`                // 杠杆倍数
	OrdType              string `json:"ordType"`              // 订单类型
	PosSide              string `json:"posSide"`              // 持仓方向
	Side                 string `json:"side"`                 // 订单方向
	SlOrdPx              string `json:"slOrdPx"`              // 止损委托价
	SlTriggerPx          string `json:"slTriggerPx"`          // 止损触发价
	SlTriggerPxType      string `json:"slTriggerPxType"`      // 止损触发价类型
	State                string `json:"state"`                // 订单状态
	Sz                   string `json:"sz"`                   // 委托数量
	Tag                  string `json:"tag"`                  // 订单标签
	TdMode               string `json:"tdMode"`               // 交易模式
	TpOrdPx              string `json:"tpOrdPx"`              // 止盈委托价
	TpTriggerPx          string `json:"tpTriggerPx"`          // 止盈触发价
	TpTriggerPxType      string `json:"tpTriggerPxType"`      // 止盈触发价类型
	UTime                string `json:"uTime"`                // 订单更新时间
}

// AlgoOrdersPendingResponse 获取未完成策略委托单列表响应
type AlgoOrdersPendingResponse struct {
	Code string                  `json:"code"`
	Msg  string                  `json:"msg"`
	Data []AlgoOrdersPendingData `json:"data"`
}

// CancelAlgoOrderRequest 撤销策略委托订单请求
type CancelAlgoOrderRequest struct {
	AlgoId string `json:"algoId"` // 策略委托单ID
	InstId string `json:"instId"` // 产品ID
}

// CancelAlgoOrdersData 撤销策略委托订单响应数据
type CancelAlgoOrdersData struct {
	AlgoClOrdId string `json:"algoClOrdId"` // 客户自定义策略订单ID
	AlgoId      string `json:"algoId"`      // 策略委托单ID
	SCode       string `json:"sCode"`       // 事件执行结果的code
	SMsg        string `json:"sMsg"`        // 事件执行失败时的msg
}

// CancelAlgoOrdersResponse 撤销策略委托订单响应
type CancelAlgoOrdersResponse struct {
	Code string                 `json:"code"`
	Msg  string                 `json:"msg"`
	Data []CancelAlgoOrdersData `json:"data"`
}

// InstrumentData 交易产品基础信息数据
type InstrumentData struct {
	InstType          string   `json:"instType"`          // 产品类型
	InstId            string   `json:"instId"`            // 产品ID
	Uly               string   `json:"uly"`               // 标的指数
	InstFamily        string   `json:"instFamily"`        // 交易品种
	Category          string   `json:"category"`          // 币种类别（已废弃）
	BaseCcy           string   `json:"baseCcy"`           // 交易货币币种
	QuoteCcy          string   `json:"quoteCcy"`          // 计价货币币种
	SettleCcy         string   `json:"settleCcy"`         // 盈亏结算和保证金币种
	CtVal             string   `json:"ctVal"`             // 合约面值
	CtMult            string   `json:"ctMult"`            // 合约乘数
	CtValCcy          string   `json:"ctValCcy"`          // 合约面值计价币种
	OptType           string   `json:"optType"`           // 期权类型
	Stk               string   `json:"stk"`               // 行权价格
	ListTime          string   `json:"listTime"`          // 上线时间
	AuctionEndTime    string   `json:"auctionEndTime"`    // 集合竞价结束时间（已废弃）
	ContTdSwTime      string   `json:"contTdSwTime"`      // 连续交易开始时间
	OpenType          string   `json:"openType"`          // 开盘类型
	ExpTime           string   `json:"expTime"`           // 产品下线时间
	Lever             string   `json:"lever"`             // 最大杠杆倍数
	TickSz            string   `json:"tickSz"`            // 下单价格精度
	LotSz             string   `json:"lotSz"`             // 下单数量精度
	MinSz             string   `json:"minSz"`             // 最小下单数量
	CtType            string   `json:"ctType"`            // 合约类型
	Alias             string   `json:"alias"`             // 合约日期别名
	State             string   `json:"state"`             // 产品状态
	RuleType          string   `json:"ruleType"`          // 交易规则类型
	MaxLmtSz          string   `json:"maxLmtSz"`          // 限价单的单笔最大委托数量
	MaxMktSz          string   `json:"maxMktSz"`          // 市价单的单笔最大委托数量
	MaxLmtAmt         string   `json:"maxLmtAmt"`         // 限价单的单笔最大美元价值
	MaxMktAmt         string   `json:"maxMktAmt"`         // 市价单的单笔最大美元价值
	MaxTwapSz         string   `json:"maxTwapSz"`         // 时间加权单的单笔最大委托数量
	MaxIcebergSz      string   `json:"maxIcebergSz"`      // 冰山委托的单笔最大委托数量
	MaxTriggerSz      string   `json:"maxTriggerSz"`      // 计划委托的单笔最大委托数量
	MaxStopSz         string   `json:"maxStopSz"`         // 止盈止损市价委托的单笔最大委托数量
	FutureSettlement  bool     `json:"futureSettlement"`  // 交割合约是否支持每日结算
	TradeQuoteCcyList []string `json:"tradeQuoteCcyList"` // 可用于交易的计价币种列表
}

// InstrumentsResponse 获取交易产品基础信息响应
type InstrumentsResponse struct {
	Code string           `json:"code"`
	Msg  string           `json:"msg"`
	Data []InstrumentData `json:"data"`
}

// SavingsBalanceData 余币宝余额数据
type SavingsBalanceData struct {
	Ccy        string `json:"ccy"`        // 币种，如 BTC
	Amt        string `json:"amt"`        // 币种数量
	Earnings   string `json:"earnings"`   // 币种持仓收益
	Rate       string `json:"rate"`       // 最新出借利率
	LoanAmt    string `json:"loanAmt"`    // 已出借数量
	PendingAmt string `json:"pendingAmt"` // 未出借数量
	RedemptAmt string `json:"redemptAmt"` // 赎回中的数量（已废弃）
}

// SavingsBalanceResponse 获取余币宝余额响应
type SavingsBalanceResponse struct {
	Code string               `json:"code"`
	Msg  string               `json:"msg"`
	Data []SavingsBalanceData `json:"data"`
}

// SavingsPurchaseRedemptRequest 余币宝申购/赎回请求
type SavingsPurchaseRedemptRequest struct {
	Ccy  string `json:"ccy"`            // 币种名称，如 BTC
	Amt  string `json:"amt"`            // 申购/赎回 数量
	Side string `json:"side"`           // 操作类型 purchase：申购 redempt：赎回
	Rate string `json:"rate,omitempty"` // 申购年利率，如 0.1代表10%，仅适用于申购
}

// SavingsPurchaseRedemptData 余币宝申购/赎回响应数据
type SavingsPurchaseRedemptData struct {
	Ccy  string `json:"ccy"`  // 币种名称
	Amt  string `json:"amt"`  // 申购/赎回 数量
	Side string `json:"side"` // 操作类型
	Rate string `json:"rate"` // 申购年利率，如 0.1代表10%
}

// SavingsPurchaseRedemptResponse 余币宝申购/赎回响应
type SavingsPurchaseRedemptResponse struct {
	Code string                       `json:"code"`
	Msg  string                       `json:"msg"`
	Data []SavingsPurchaseRedemptData `json:"data"`
}

// AssetBalanceData 资金账户余额数据
type AssetBalanceData struct {
	Ccy       string `json:"ccy"`       // 币种，如 BTC
	Bal       string `json:"bal"`       // 余额
	FrozenBal string `json:"frozenBal"` // 冻结余额
	AvailBal  string `json:"availBal"`  // 可用余额
}

// AssetBalanceResponse 获取资金账户余额响应
type AssetBalanceResponse struct {
	Code string             `json:"code"`
	Msg  string             `json:"msg"`
	Data []AssetBalanceData `json:"data"`
}

// AssetTransferRequest 资金划转请求
type AssetTransferRequest struct {
	Type        string `json:"type,omitempty"`        // 划转类型，默认为0（账户内划转）
	Ccy         string `json:"ccy"`                   // 划转币种，如 USDT
	Amt         string `json:"amt"`                   // 划转数量
	From        string `json:"from"`                  // 转出账户 6：资金账户 18：交易账户
	To          string `json:"to"`                    // 转入账户 6：资金账户 18：交易账户
	SubAcct     string `json:"subAcct,omitempty"`     // 子账户名称
	LoanTrans   bool   `json:"loanTrans,omitempty"`   // 是否支持借币转出，默认为false
	OmitPosRisk string `json:"omitPosRisk,omitempty"` // 是否忽略仓位风险，默认为false
	ClientId    string `json:"clientId,omitempty"`    // 客户自定义ID
}

// AssetTransferData 资金划转响应数据
type AssetTransferData struct {
	TransId  string `json:"transId"`  // 划转ID
	Ccy      string `json:"ccy"`      // 划转币种
	ClientId string `json:"clientId"` // 客户自定义ID
	From     string `json:"from"`     // 转出账户
	Amt      string `json:"amt"`      // 划转量
	To       string `json:"to"`       // 转入账户
}

// AssetTransferResponse 资金划转响应
type AssetTransferResponse struct {
	Code string              `json:"code"`
	Msg  string              `json:"msg"`
	Data []AssetTransferData `json:"data"`
}

// KlineResponse 是OKX API响应的结构
type HttpKlineResponse struct {
	Code string     `json:"code"`
	Msg  string     `json:"msg"`
	Data [][]string `json:"data"`
}
