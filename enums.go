package galatvtr

// 交易所枚举
type Exchange string

const (
	// 主要交易所
	ExchangeOKX  Exchange = "OKX"
	ExchangeGate Exchange = "GATEIO"
)

// 获取交易所显示名称
func (e Exchange) DisplayName() string {
	switch e {
	case ExchangeOKX:
		return "OKX"
	case ExchangeGate:
		return "GATEIO"
	default:
		return string(e)
	}
}
